package java

import (
	"github.com/aadog/go-jni"
	"runtime"
	"syscall"
)

/*
#include <pthread.h>
#include <stdlib.h>
extern void* PerformThreadFunc(void*);
*/
import "C"

var Jvm jni.VM

//func JvmPatchGlobal() {
//	globals := unsafe.Add(Jvm, 112)
//
//	ndk.LogError("Go", "globals:%v", globals)
//	ndk.LogError("Go", "max:%v", ffi.Ptr(unsafe.Add(globals, 48)).ReadInt())
//	ndk.LogError("Go", "globals DumpHex:%v", hex.Dump(ffi.Ptr(globals).ReadByteArray(500)))
//}
//func SetJvm(ptr unsafe.Pointer) {
//	if uintptr(ptr) == 0 {
//		return
//	}
//	Jvm = jni.VM(ptr)
//	env, _ := Jvm.GetEnv(jni.JNI_VERSION_1_6)
//	if env == 0 {
//		env, _ = Jvm.AttachCurrentThreadAsDaemon()
//	}
//	//go gcThreadLoop()
//}

func Perform(fn func()) chan struct{} {
	done := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		SetLocalThreadJavaEnv()
		defer RemoveLocalThreadJavaEnv()
		LogInfo("Go", "start thread 线程id:%v", syscall.Gettid())
		fn()
		done <- struct{}{}
	}()
	return done
}
