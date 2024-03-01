package java

//
//var Modifier JavaLangReflectModifier
//
//type JavaLangReflectModifier struct {
//}
//
//func (j JavaLangReflectModifier) IsAbstract(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isAbstract", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsFinal(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isFinal", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsInterface(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isInterface", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsPrivate(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isPrivate", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsNative(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isNative", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsProtected(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isProtected", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsPublic(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java.lang.reflect.Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isPublic", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsStatic(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isStatic", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsStrict(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isStrict", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsSynchronized(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isSynchronized", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsTransient(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isTransient", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//func (j JavaLangReflectModifier) IsVolatile(modifiers int) mo.Result[bool] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	isStaticMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "isVolatile", "(I)Z").Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	defer env.DeleteLocalRef(isStaticMethod)
//	b := env.CallStaticBooleanMethodA(modifierClass.JniPtr(), isStaticMethod, jni.Jvalue(modifiers))
//	return b
//}
//
//func (j JavaLangReflectModifier) ToString(modifiers int) mo.Result[string] {
//	env := LocalThreadJavaEnv()
//	modifierClass, err := Use("java/lang/reflect/Modifier").Get()
//	if err != nil {
//		return mo.Err[string](err)
//	}
//	toStringMethod, err := env.GetStaticMethodID(modifierClass.JniPtr(), "toString", "(I)Ljava/lang/String;").Get()
//	if err != nil {
//		return mo.Err[string](err)
//	}
//	defer env.DeleteLocalRef(toStringMethod)
//	methodSignatureString, err := env.CallStaticObjectMethodA(modifierClass.JniPtr(), toStringMethod).Get()
//	if err != nil {
//		return mo.Err[string](err)
//	}
//	defer env.DeleteLocalRef(methodSignatureString)
//	methodSignature := string(env.GetStringUTF(methodSignatureString))
//	return mo.Ok(methodSignature)
//}
