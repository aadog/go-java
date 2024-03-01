package java

//
//type JavaLangReflectConstructorObjectWrapper struct {
//	*Object
//}
//
//func (j *JavaLangReflectConstructorObjectWrapper) GetModifiers() mo.Result[int] {
//	env := LocalThreadJavaEnv()
//	cls, err := j.Class().Get()
//	if err != nil {
//		return mo.Err[int](err)
//	}
//	getParameterCountMethodId, err := cls.GetMethodID("getModifiers", "()I").Get()
//	if err != nil {
//		return mo.Err[int](err)
//	}
//	return env.CallIntMethodA(j.JniPtr(), getParameterCountMethodId)
//}
//func (j *JavaLangReflectConstructorObjectWrapper) GetParameterCount() mo.Result[int] {
//	env := LocalThreadJavaEnv()
//	objCls, err := j.Class().Get()
//	if err != nil {
//		return mo.Err[int](err)
//	}
//	getParameterCountMethodId, err := objCls.GetMethodID("getParameterCount", "()I").Get()
//	if err != nil {
//		return mo.Err[int](err)
//	}
//	return env.CallIntMethodA(j.JniPtr(), getParameterCountMethodId)
//}
//
////	func (j *JavaLangReflectConstructorObjectWrapper) ToGenericString() mo.Result[string] {
////		objCls, err := j.ObjectClass().Get()
////		if err != nil {
////			return mo.Err[string](err)
////		}
////		getParameterCountMethodId, err := objCls.GetMethodId("toGenericString", "()Ljava/lang/String;").Get()
////		if err != nil {
////			return mo.Err[string](err)
////		}
////		return j.CallStringMethodA(getParameterCountMethodId)
////	}
//func (j *JavaLangReflectConstructorObjectWrapper) GetParameterTypes() mo.Result[[]*Class] {
//	env := LocalThreadJavaEnv()
//	objCls, err := j.Class().Get()
//	if err != nil {
//		return mo.Err[[]*Class](err)
//	}
//	getParameterTypesMethodId, err := objCls.GetMethodID("getParameterTypes", "()[Ljava/lang/Class;").Get()
//	if err != nil {
//		return mo.Err[[]*Class](err)
//	}
//
//	parameterTypeArray, err := env.CallObjectMethodA(j.JniPtr(), getParameterTypesMethodId).Get()
//	if err != nil {
//		return mo.Err[[]*Class](err)
//	}
//	defer env.DeleteLocalRef(parameterTypeArray)
//
//	n := env.GetArrayLength(parameterTypeArray)
//	clsWrappers := make([]*Class, 0)
//	for i := 0; i < n; i++ {
//		item := env.GetObjectArrayElement(parameterTypeArray, i)
//		if item == 0 {
//			panic(errors.New("发生错误"))
//		}
//		defer env.DeleteLocalRef(item)
//
//		clsName, err := GetClassName(item).Get()
//		if err != nil {
//			return mo.Err[[]*Class](err)
//		}
//		_ = clsName
//		cls, err := Use(clsName).Get()
//		if err != nil {
//			return mo.Err[[]*Class](err)
//		}
//
//		clsWrappers = append(clsWrappers, cls)
//	}
//	return mo.Ok(clsWrappers)
//}
//
//func JavaLangReflectConstructorWithJniPtr(ptr uintptr) *JavaLangReflectConstructorObjectWrapper {
//	return &JavaLangReflectConstructorObjectWrapper{
//		Object: ObjectWrapperWithJniPtr(ptr),
//	}
//}
