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
	"reflect"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
)

var Envs sync.Map
var PrimitiveClasses = []string{
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
	env, ret := Jvm.AttachCurrentThread()
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
	threadID := CurrentThreadID()
	Envs.Delete(threadID)
	Jvm.DetachCurrentThread()
}

var ClassWrapperCacheMap = sync.Map{}
var lkClassWrapperMap = sync.Mutex{}

func Use(className string, skipCache bool) mo.Result[*Class] {
	className = strings.ReplaceAll(className, ".", "/")

	lkClassWrapperMap.Lock()
	defer lkClassWrapperMap.Unlock()
	if !skipCache {
		loadClass, ok := ClassWrapperCacheMap.Load(className)
		if ok {
			return mo.Ok(loadClass.(*Class))
		}
	}
	env := LocalThreadJavaEnv()
	var cls jni.Jclass
	if lo.Contains(PrimitiveClasses, className) {
		var err error
		//LogError("Go", "查找:%v", className)
		cls, err = GetPrimitiveClass(className).Get()
		//LogError("Go", "查找结果:%v", cls)
		if err != nil {
			return mo.Errf[*Class]("find class error:%v", err)
		}
		defer env.DeleteLocalRef(cls)
	} else {
		var err error
		cls, err = env.FindClass(className).Get()
		if err != nil {
			return mo.Errf[*Class]("find class error:%v", err)
		}
		defer env.DeleteLocalRef(cls)
	}
	if cls == 0 {
		if env.ExceptionCheck() {
			return mo.Errf[*Class](env.GetAndClearExceptionMessage())
		}
	}

	//所有class都使用GlobalRef,
	clsWrapper := JniObjectWithPtrAndNewGlobalRef[Class](cls)
	ClassWrapperCacheMap.Store(className, clsWrapper)
	return mo.Ok(clsWrapper)
}
func UseT[T any](className string, skipCache bool) mo.Result[*T] {
	var t T
	cls, err := Use(className, skipCache).Get()
	if err != nil {
		return mo.Err[*T](err)
	}
	_ = cls
	vl := reflect.ValueOf(&t)
	vl.Elem().FieldByName("Class").Set(reflect.ValueOf(cls))
	return mo.Ok(&t)
}
func StoreClass(className string, cls jni.Jclass) {
	className = strings.ReplaceAll(className, ".", "/")
	lkClassWrapperMap.Lock()
	defer lkClassWrapperMap.Unlock()
	_, ok := ClassWrapperCacheMap.Load(className)
	if ok {
		return
	}
	//所有class都使用GlobalRef,
	clsWrapper := JniObjectWithPtrAndNewGlobalRef[Class](cls)
	ClassWrapperCacheMap.Store(className, clsWrapper)
}
