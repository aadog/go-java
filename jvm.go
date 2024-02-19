package java

import (
	"runtime"
)

func Perform(fn func()) chan struct{} {
	done := make(chan struct{})
	go func() {
		runtime.LockOSThread()
		SetLocalThreadJavaEnv()
		defer RemoveLocalThreadJavaEnv()
		fn()
		done <- struct{}{}
	}()
	return done
}
