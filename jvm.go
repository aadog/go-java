package java

import (
	"runtime"
	"syscall"
)

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
