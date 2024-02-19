package java

/*
#include <pthread.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/aadog/go-jni"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
)

var Envs sync.Map
var PrimitiveClasss = []string{
	"byte", "short", "int", "long", "float", "double", "boolean", "char",
}

func CurrentThreadID() uint64 {
	return uint64(syscall.Gettid())
}

func LocalThreadJavaEnv() jni.Env {
	threadID := CurrentThreadID()
	env, ok := Envs.Load(threadID)
	if !ok {
		LogInfo("Go", fmt.Sprintf("threadID:%d", threadID))
		panic(errors.New(fmt.Sprintf("获取java env失败1:%s", string(debug.Stack()))))
	}
	if env == 0 {
		panic(errors.New(fmt.Sprintf("获取java env失败2:%s", string(debug.Stack()))))
	}
	return env.(jni.Env)
}
func SetLocalThreadJavaEnv() {
	env, ret := Android.Jvm.AttachCurrentThread()
	if ret != 0 {
		panic(errors.New("attach current thread error"))
	}
	if env == 0 {
		panic(errors.New(fmt.Sprintf("设置java env失败:%s", string(debug.Stack()))))
	}
	threadID := CurrentThreadID()
	Envs.Store(threadID, env)
}
func RemoveLocalThreadJavaEnv() {
	LogInfo("Go", "线程结束:%v", "111")
	threadID := CurrentThreadID()
	Envs.Delete(threadID)
	Android.Jvm.DetachCurrentThread()
	LogInfo("Go", "线程结束:%v", threadID)
}

var ClassWrapperCacheMap = sync.Map{}
var lkClassWrapperMap = sync.Mutex{}

func Use(className string) mo.Result[*ClassWrapper] {
	className = strings.ReplaceAll(className, ".", "/")
	lkClassWrapperMap.Lock()
	defer lkClassWrapperMap.Unlock()
	loadClass, ok := ClassWrapperCacheMap.Load(className)
	if ok {
		return mo.Ok(loadClass.(*ClassWrapper))
	}
	env := LocalThreadJavaEnv()

	var cls jni.Jclass

	if lo.Contains(PrimitiveClasss, className) {
		var err error
		LogError("Go", "查找:%v", className)
		cls, err = GetPrimitiveClass(className).Get()
		LogError("Go", "查找结果:%v", cls)
		if err != nil {
			return mo.Errf[*ClassWrapper]("find class error:%v", err)
		}
		defer env.DeleteLocalRef(cls)
	} else {
		var err error
		cls, err = env.FindClass(className).Get()
		if err != nil {
			return mo.Errf[*ClassWrapper]("find class error:%v", err)
		}
		defer env.DeleteLocalRef(cls)
	}
	if cls == 0 {
		if env.ExceptionCheck() {
			return mo.Errf[*ClassWrapper](env.GetAndClearExceptionMessage())
		}
	}

	//所有class都使用GlobalRef,
	clsWrapper := ClassWrapperWithJniPtr(env.NewGlobalRef(cls))
	ClassWrapperCacheMap.Store(className, clsWrapper)
	return mo.Ok(clsWrapper)
}
func ForeUse(className string, cls jni.Jclass) mo.Result[*ClassWrapper] {
	className = strings.ReplaceAll(className, ".", "/")
	lkClassWrapperMap.Lock()
	defer lkClassWrapperMap.Unlock()
	loadClass, ok := ClassWrapperCacheMap.Load(className)
	if ok {
		return mo.Ok(loadClass.(*ClassWrapper))
	}
	env := LocalThreadJavaEnv()

	//所有class都使用GlobalRef,
	clsWrapper := ClassWrapperWithJniPtr(env.NewGlobalRef(cls))
	ClassWrapperCacheMap.Store(className, clsWrapper)
	return mo.Ok(clsWrapper)
}
