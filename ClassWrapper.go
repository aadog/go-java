package java

import "C"
import (
	"errors"
	"fmt"
	"github.com/aadog/go-jni"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"strings"
)

type ClassWrapper struct {
	ptr uintptr
}

func (c *ClassWrapper) GetMethodID(name string, sig string) mo.Result[jni.JmethodID] {
	env := LocalThreadJavaEnv()
	method, err := env.GetMethodID(c.JniPtr(), name, sig).Get()
	if err != nil {
		return mo.Err[jni.JmethodID](err)
	}
	return mo.Ok(method)
}

//	func (c *ClassWrapper) MatchTypes(types []*ClassWrapper, sType []string) mo.Result[bool] {
//		//LogInfo("Go", "444")
//		for idx, wrapper := range types {
//			//LogInfo("Go", "555")
//			typeName, err := wrapper.GetName().Get()
//			if err != nil {
//				return mo.Err[bool](err)
//			}
//			//LogInfo("Go", "777")
//			LogInfo("Go", "对比%v=%v", typeName, sType[idx])
//			if typeName != sType[idx] {
//				return mo.Ok(false)
//			}
//		}
//		//LogInfo("Go", "666")
//		return mo.Ok(true)
//	}

func (c *ClassWrapper) EnumMatchStaticMethod(funcName string, classTypeName ...string) mo.Result[*JavaLangReflectMethodObjectWrapper] {
	allMethods, err := c.GetMethods().Get()
	if err != nil {
		return mo.Err[*JavaLangReflectMethodObjectWrapper](err)
	}
	var matchPtr uintptr
	defer func() {
		for _, method := range allMethods {
			if method.ptr != matchPtr {
				//LogError("Go", "free:%v", funcName, method.ptr)
				method.Free()
			}
		}
	}()

	match, finded := lo.Find(allMethods, func(item *JavaLangReflectMethodObjectWrapper) bool {
		modifiers := item.GetModifiers().MustGet()
		if !Modifier.IsPublic(modifiers).MustGet() {
			return false
		}
		if !Modifier.IsStatic(modifiers).MustGet() {
			return false
		}
		itemName := item.GetName().MustGet()
		if itemName != funcName {
			return false
		}

		paramCount := item.GetParameterCount().MustGet()
		if paramCount != len(classTypeName) {
			return false
		}
		itemTypes, err := item.GetParameterTypes().Get()
		if err != nil {
			panic(err)
		}
		for idx, itemType := range itemTypes {
			cls := Use(classTypeName[idx]).MustGet()
			if !cls.IsAssignableFrom(itemType) {
				return false
			}
		}
		return true
	})
	if finded == false {
		sErrBuilder := strings.Builder{}
		sErrBuilder.WriteString(fmt.Sprintf("not match static %s methods, look method list\n", funcName))
		sErrBuilder.WriteString(DeclaredMethodsToString(lo.Filter(allMethods, func(item *JavaLangReflectMethodObjectWrapper, index int) bool {
			modifiers := item.GetModifiers().MustGet()
			if !Modifier.IsPublic(modifiers).MustGet() {
				return false
			}
			if !Modifier.IsStatic(modifiers).MustGet() {
				return false
			}
			itemName := item.GetName().MustGet()
			if itemName != funcName {
				return false
			}
			return true
		})))

		return mo.Err[*JavaLangReflectMethodObjectWrapper](errors.New(sErrBuilder.String()))
	}
	matchPtr = match.ptr
	return mo.Ok(match)
}
func (c *ClassWrapper) EnumMatchMethod(funcName string, classTypeName ...string) mo.Result[*JavaLangReflectMethodObjectWrapper] {
	allMethods, err := c.GetMethods().Get()
	if err != nil {
		return mo.Err[*JavaLangReflectMethodObjectWrapper](err)
	}
	var matchPtr uintptr
	defer func() {
		for _, method := range allMethods {
			if method.ptr != matchPtr {
				//LogError("Go", "free:%v", funcName, method.ptr)
				method.Free()
			}
		}
	}()

	match, finded := lo.Find(allMethods, func(item *JavaLangReflectMethodObjectWrapper) bool {
		modifiers := item.GetModifiers().MustGet()
		if !Modifier.IsPublic(modifiers).MustGet() {
			return false
		}
		if Modifier.IsStatic(modifiers).MustGet() {
			return false
		}
		itemName := item.GetName().MustGet()
		if itemName != funcName {
			return false
		}

		paramCount := item.GetParameterCount().MustGet()
		if paramCount != len(classTypeName) {
			return false
		}
		itemTypes, err := item.GetParameterTypes().Get()
		if err != nil {
			panic(err)
		}
		if paramCount != len(classTypeName) {
			return false
		}
		for idx, itemType := range itemTypes {
			cls := Use(classTypeName[idx]).MustGet()
			if !cls.IsAssignableFrom(itemType) {
				return false
			}
		}
		return true
	})
	if finded == false {
		sErrBuilder := strings.Builder{}
		sErrBuilder.WriteString(fmt.Sprintf("not match %s methods, look method list\n", funcName))
		sErrBuilder.WriteString(DeclaredMethodsToString(lo.Filter(allMethods, func(item *JavaLangReflectMethodObjectWrapper, index int) bool {
			modifiers := item.GetModifiers().MustGet()
			if !Modifier.IsPublic(modifiers).MustGet() {
				return false
			}
			if Modifier.IsStatic(modifiers).MustGet() {
				return false
			}
			itemName := item.GetName().MustGet()
			if itemName != funcName {
				return false
			}
			return true
		})))

		return mo.Err[*JavaLangReflectMethodObjectWrapper](errors.New(sErrBuilder.String()))
	}
	matchPtr = match.ptr
	return mo.Ok(match)
}
func (c *ClassWrapper) MatchMethod(funcName string, classTypeName ...string) mo.Result[*JavaLangReflectMethodObjectWrapper] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[*JavaLangReflectMethodObjectWrapper](err)
	}
	getConstructorMethodId, err := objCls.GetMethodID("getMethod", "(Ljava/lang/String;[Ljava/lang/Class;)Ljava/lang/reflect/Method;").Get()
	if err != nil {
		return mo.Err[*JavaLangReflectMethodObjectWrapper](err)
	}

	jarray := env.NewObjectArray(len(classTypeName), objCls.JniPtr(), 0)
	defer env.DeleteLocalRef(jarray)
	for i, s := range classTypeName {
		env.SetObjectArrayElement(jarray, i, Use(s).MustGet().JniPtr())
	}

	methodName := env.NewString(funcName)
	defer env.DeleteLocalRef(methodName)
	construstor, err := env.CallObjectMethodA(c.JniPtr(), getConstructorMethodId, jni.Jvalue(methodName), jni.Jvalue(jarray)).Get()
	if err != nil {
		allConstructors, err := c.GetMethods().Get()
		if err != nil {
			panic(err)
		}
		defer func() {
			for _, constructor := range allConstructors {
				constructor.Free()
			}
		}()
		filterMethods := lo.Filter(allConstructors, func(item *JavaLangReflectMethodObjectWrapper, index int) bool {
			if item.GetName().OrElse("") == funcName {
				return true
			}
			return false
		})
		sErrBuilder := strings.Builder{}
		sErrBuilder.WriteString(fmt.Sprintf("not match %s method, look method list\n", funcName))
		sErrBuilder.WriteString(DeclaredMethodsToString(filterMethods))

		return mo.Err[*JavaLangReflectMethodObjectWrapper](errors.New(sErrBuilder.String()))
	}
	return mo.Ok(JavaLangReflectMethodWithJniPtr(construstor))
}
func (c *ClassWrapper) MatchConstructor(classTypeName ...string) mo.Result[*JavaLangReflectConstructorObjectWrapper] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[*JavaLangReflectConstructorObjectWrapper](err)
	}
	getConstructorMethodId, err := objCls.GetMethodID("getConstructor", "([Ljava/lang/Class;)Ljava/lang/reflect/Constructor;").Get()
	if err != nil {
		return mo.Err[*JavaLangReflectConstructorObjectWrapper](err)
	}

	jarray := env.NewObjectArray(len(classTypeName), objCls.JniPtr(), 0)
	defer env.DeleteLocalRef(jarray)
	for i, s := range classTypeName {
		env.SetObjectArrayElement(jarray, i, Use(s).MustGet().JniPtr())
	}

	construstor, err := env.CallObjectMethodA(c.JniPtr(), getConstructorMethodId, jni.Jvalue(jarray)).Get()
	if err != nil {
		allConstructors, err := c.GetConstructors().Get()
		if err != nil {
			panic(err)
		}
		defer func() {
			for _, constructor := range allConstructors {
				constructor.Free()
			}
		}()
		sErrBuilder := strings.Builder{}
		sErrBuilder.WriteString(fmt.Sprintf("not match %s constructor, look constructor list\n", c.GetName().MustGet()))
		sErrBuilder.WriteString(DeclaredConstructorsToString(allConstructors))

		return mo.Err[*JavaLangReflectConstructorObjectWrapper](errors.New(sErrBuilder.String()))
	}
	return mo.Ok(JavaLangReflectConstructorWithJniPtr(construstor))
}
func (c *ClassWrapper) GetMethods() mo.Result[[]*JavaLangReflectMethodObjectWrapper] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectMethodObjectWrapper](err)
	}
	getConstructorsMethodId, err := env.GetMethodID(objCls.JniPtr(), "getMethods", "()[Ljava/lang/reflect/Method;").Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectMethodObjectWrapper](err)
	}
	constructorsArray, err := env.CallObjectMethodA(c.JniPtr(), getConstructorsMethodId).Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectMethodObjectWrapper](err)
	}
	defer env.DeleteLocalRef(constructorsArray)

	constructorsCount := env.GetArrayLength(constructorsArray)

	allMethod := make([]*JavaLangReflectMethodObjectWrapper, 0)

	for i := 0; i < constructorsCount; i++ {
		methodObj := env.GetObjectArrayElement(constructorsArray, i)
		allMethod = append(allMethod, JavaLangReflectMethodWithJniPtr(methodObj))
	}
	return mo.Ok(allMethod)
}
func (c *ClassWrapper) GetConstructors() mo.Result[[]*JavaLangReflectConstructorObjectWrapper] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectConstructorObjectWrapper](err)
	}
	getConstructorsMethodId, err := env.GetMethodID(objCls.JniPtr(), "getConstructors", "()[Ljava/lang/reflect/Constructor;").Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectConstructorObjectWrapper](err)
	}
	constructorsArray, err := env.CallObjectMethodA(c.JniPtr(), getConstructorsMethodId).Get()
	if err != nil {
		return mo.Err[[]*JavaLangReflectConstructorObjectWrapper](err)
	}
	defer env.DeleteLocalRef(constructorsArray)

	constructorsCount := env.GetArrayLength(constructorsArray)

	allMethod := make([]*JavaLangReflectConstructorObjectWrapper, 0)

	for i := 0; i < constructorsCount; i++ {
		methodObj := env.GetObjectArrayElement(constructorsArray, i)
		wrapper := JavaLangReflectConstructorWithJniPtr(methodObj)
		allMethod = append(allMethod, wrapper)
	}
	return mo.Ok(allMethod)
}
func (c *ClassWrapper) IsAssignableFrom(cls *ClassWrapper) bool {
	env := LocalThreadJavaEnv()
	return env.IsAssignableFrom(c.JniPtr(), cls.JniPtr())
}

func (c *ClassWrapper) EnumMatchConstructor(classTypeName ...string) mo.Result[*JavaLangReflectConstructorObjectWrapper] {
	allConstructors, err := c.GetConstructors().Get()
	if err != nil {
		return mo.Err[*JavaLangReflectConstructorObjectWrapper](err)
	}

	var matchPtr uintptr
	defer func() {
		for _, constructor := range allConstructors {
			if constructor.ptr != matchPtr {
				constructor.Free()
			}
		}
	}()

	match, finded := lo.Find(allConstructors, func(item *JavaLangReflectConstructorObjectWrapper) bool {
		modifiers := item.GetModifiers().MustGet()
		if !Modifier.IsPublic(modifiers).MustGet() {
			return false
		}

		paramCount := item.GetParameterCount().MustGet()
		if paramCount != len(classTypeName) {
			return false
		}
		itemTypes := item.GetParameterTypes().MustGet()
		for idx, itemType := range itemTypes {
			cls := Use(classTypeName[idx]).MustGet()
			if !cls.IsAssignableFrom(itemType) {
				return false
			}
		}
		return true
	})
	if finded == false {
		sErrBuilder := strings.Builder{}
		sErrBuilder.WriteString(fmt.Sprintf("not match %s constructor, look constructor list\n", c.GetName().MustGet()))
		sErrBuilder.WriteString(DeclaredConstructorsToString(allConstructors))

		return mo.Err[*JavaLangReflectConstructorObjectWrapper](errors.New(sErrBuilder.String()))
	}
	matchPtr = match.ptr
	return mo.Ok(match)
}
func (c *ClassWrapper) New(args ...any) mo.Result[*ObjectWrapper] {
	constructor, err := c.EnumMatchConstructor(SumGoArgsType(args...)...).Get()
	if err != nil {
		return mo.Err[*ObjectWrapper](err)
	}
	defer constructor.Free()

	env := LocalThreadJavaEnv()
	constructorMethodId, err := env.FromReflectedMethod(constructor.JniPtr()).Get()
	if err != nil {
		return mo.Err[*ObjectWrapper](err)
	}

	jArgs := make([]jni.Jvalue, 0)
	for _, arg := range args {
		jval, needDelete := ConvertAnyArgToJValueArg(arg)
		if needDelete {
			defer env.DeleteLocalRef(jni.Jobject(jval))
		}
		jArgs = append(jArgs, jval)
	}
	obj, err := env.NewObjectA(c.JniPtr(), constructorMethodId, jArgs...).Get()
	if err != nil {
		return mo.Err[*ObjectWrapper](err)
	}
	return mo.Ok[*ObjectWrapper](ObjectWrapperWithJniPtr(obj))
}
func (c *ClassWrapper) IsArray() mo.Result[bool] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	isArrayMethodId, err := env.GetMethodID(objCls.JniPtr(), "isArray", "()Z").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	isArray, err := env.CallBooleanMethodA(c.JniPtr(), isArrayMethodId).Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(isArray)
}
func (c *ClassWrapper) IsInnerClass() mo.Result[bool] {
	env := LocalThreadJavaEnv()
	cls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(env.IsSameObject(c.JniPtr(), cls.JniPtr()))
}

func (c *ClassWrapper) String() string {
	s, err := c.ToString().Get()
	if err != nil {
		panic(err)
	}
	return s
}
func (o *ClassWrapper) CallStaticObjectA(funcName string, args ...any) mo.Result[*ObjectWrapper] {
	ret, err := ClassWrapperStaticCall[*ObjectWrapper](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[*ObjectWrapper](err)
	}
	return mo.Ok(ret.(*ObjectWrapper))
}
func (o *ClassWrapper) CallStaticStringA(funcName string, args ...any) mo.Result[string] {
	ret, err := ClassWrapperStaticCall[string](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(ret.(string))
}

//	func (c *ClassWrapper) ToString() mo.Result[string] {
//		toStringMethodId, err := c.GetMethodId("toString", "()Ljava/lang/String;").Get()
//		if err != nil {
//			return mo.Err[string](err)
//		}
//		jstring, err := c.CallStringMethodA(toStringMethodId).Get()
//		if err != nil {
//			return mo.Err[string](err)
//		}
//		return mo.Ok(jstring)
//	}

func (c *ClassWrapper) GetSuperclass() mo.Result[*ClassWrapper] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[*ClassWrapper](err)
	}
	getNameMethodId, err := env.GetMethodID(objCls.JniPtr(), "getSuperclass", "()Ljava/lang/Class;").Get()
	if err != nil {
		return mo.Err[*ClassWrapper](err)
	}
	jstringName, err := env.CallObjectMethodA(c.ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[*ClassWrapper](err)
	}
	defer env.DeleteLocalRef(jstringName)
	clsName, err := GetClassName(jstringName).Get()
	if err != nil {
		return mo.Err[*ClassWrapper](err)
	}
	return Use(clsName)
}
func (c *ClassWrapper) GetSimpleName() mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethodId, err := env.GetMethodID(objCls.JniPtr(), "getSimpleName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(c.ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}
func (c *ClassWrapper) GetName() mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethodId, err := env.GetMethodID(objCls.JniPtr(), "getName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(c.ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}
func (c *ClassWrapper) ToString() mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethodId, err := env.GetMethodID(objCls.JniPtr(), "toString", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(c.ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}
func (c *ClassWrapper) ToGenericString() mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethodId, err := env.GetMethodID(objCls.JniPtr(), "toGenericString", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(c.ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}
func (c *ClassWrapper) JniPtr() uintptr {
	return c.ptr
}

//	func (c *ClassWrapper) Free() {
//		//class no need free
//		//env := LocalThreadJavaEnv()
//		//env.DeleteGlobalRef(c.ptr)
//	}
func ClassWrapperWithJniPtr(ptr uintptr) *ClassWrapper {
	cls := &ClassWrapper{ptr: ptr}
	return cls
}
