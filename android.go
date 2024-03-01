package java

/*
#include <dlfcn.h>
*/
import "C"
import (
	"errors"
	"github.com/aadog/go-ffi"
	"github.com/aadog/go-jni"
	"github.com/samber/mo"
	"runtime"
	"unsafe"
)

var Jvm jni.VM

func FindJNI_GetCreatedJavaVMs() ffi.NativePointer {
	handle := C.dlopen(nil, C.RTLD_NOW)
	JNI_GetCreatedJavaVMsPtr := C.dlsym(handle, (*C.char)(jni.CString("JNI_GetCreatedJavaVMs")))
	C.dlclose(handle)
	return ffi.Ptr(uintptr(JNI_GetCreatedJavaVMsPtr))
}

// InitWithArtFind 初始化jvm
func InitWithArtFind() mo.Result[struct{}] {
	//JNI_GetCreatedJavaVMsPtr := gumjs.Module.FindSymbolByName(lo.ToPtr("libart.so"), "JNI_GetCreatedJavaVMs")
	JNI_GetCreatedJavaVMsPtr := FindJNI_GetCreatedJavaVMs()
	if JNI_GetCreatedJavaVMsPtr.IsNull() {
		return mo.Err[struct{}](errors.New("art find JNI_GetCreatedJavaVMs error"))
	}
	JNI_GetCreatedJavaVMs := ffi.NewNativeFunction(JNI_GetCreatedJavaVMsPtr.ToUinptr(), ffi.Tint, []ffi.ArgTypeName{ffi.TPointer, ffi.Tint, ffi.TPointer})
	var vms *unsafe.Pointer
	var vmCount int
	JNI_GetCreatedJavaVMs.Call(unsafe.Pointer(&vms), 1, &vmCount)
	if vmCount < 1 {
		return mo.Err[struct{}](errors.New("vmcount <1"))
	}
	Jvm = jni.VM(unsafe.Pointer(vms))
	return mo.Ok(struct{}{})
}

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
