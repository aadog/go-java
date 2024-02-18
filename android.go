package java

import (
	"encoding/hex"
	"errors"
	"github.com/aadog/go-ffi"
	"github.com/aadog/go-jni"
	"github.com/samber/lo"
	"unsafe"
)

var Android AndroidStruct

type AndroidStruct struct {
	Jvm                    jni.VM
	Vm                     ffi.NativePointer
	GlobalTable            ffi.NativePointer
	GlobalTableMaxEntries_ ffi.NativePointer
	GlobalLock             ffi.NativePointer
}

func (a AndroidStruct) PatchGlobalRef() {
	if unsafe.Sizeof(lo.ToPtr(1)) == 4 {
		Android.GlobalLock = Android.Vm.Add(32)
		Android.GlobalTable = Android.Vm.Add(72)
		panic(errors.New("没有计算，使用hexdump打印看一下位置"))
	} else {
		Android.GlobalLock = Android.Vm.Add(64)
		Android.GlobalTable = Android.Vm.Add(112)
		Android.GlobalTableMaxEntries_ = Android.GlobalTable.Add(3 * 16)
		LogError("Go", "%s", hex.Dump(Android.GlobalTable.ReadByteArray(500)))

	}
	//Android.GlobalTableMaxEntries_.WriteU32(0xffffffff)
	//Android.GlobalTableMaxEntries_.WriteU32(100)
	Android.GlobalTableMaxEntries_.WriteSize(3048)
	LogError("Go", "%d", Android.GlobalTableMaxEntries_.ReadU32())
	LogError("Go", "%s", hex.Dump(Android.GlobalTable.ReadByteArray(500)))
}
func (a AndroidStruct) InitWithArtFind() error {
	JNI_GetCreatedJavaVMsPtr := gumjs.Module.FindSymbolByName(lo.ToPtr("libart.so"), "JNI_GetCreatedJavaVMs")
	if JNI_GetCreatedJavaVMsPtr.IsNull() {
		return errors.New("art find JNI_GetCreatedJavaVMs error")
	}
	JNI_GetCreatedJavaVMs := ffi.NewNativeFunction(JNI_GetCreatedJavaVMsPtr.ToUinptr(), ffi.Tint, []ffi.ArgTypeName{ffi.TPointer, ffi.Tint, ffi.TPointer})
	var vms *unsafe.Pointer
	var vmCount int
	JNI_GetCreatedJavaVMs.Call(unsafe.Pointer(&vms), 1, &vmCount)
	if vmCount < 1 {
		return errors.New("vmcount <1")
	}
	Android.Vm = ffi.Ptr(uintptr(unsafe.Pointer(vms)))
	Android.Jvm = jni.VM(Android.Vm.ToUinptr())
	//Android.init()
	return nil
}

func (a AndroidStruct) init() {
	a.PatchGlobalRef()
}
