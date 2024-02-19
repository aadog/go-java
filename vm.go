package java

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
