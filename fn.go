package java

import (
	"errors"
	"github.com/aadog/go-jni"
	"github.com/samber/mo"
)

type JniCallRet interface {
	*Object | string | struct{} | int | *string | int64 | bool
}
type JniFieldRet interface {
	*Object | string | int | *string | int64 | bool
}

func ClassCallStaticMethod[T JniCallRet](o *Class, funcName string, sig string, args ...any) mo.Result[any] {
	env := LocalThreadJavaEnv()
	objCls := o
	methodId, err := objCls.GetStaticMethodID(funcName, sig).Get()
	if err != nil {
		return mo.Err[any](err)
	}
	jArgs := make([]jni.Jvalue, 0)
	for _, arg := range args {
		jval, needDelete := ConvertAnyArgToJValueArg(arg)
		if needDelete {
			defer env.DeleteLocalRef(jni.Jobject(jval))
		}
		jArgs = append(jArgs, jval)
	}

	var inputType T

	switch any(inputType).(type) {
	case *Object:
		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		objWrapper := JniObjectWithPtr[Object](obj)
		return mo.Ok(any(objWrapper))
	case *Class:
		panic(errors.New("system error"))
		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
		//if err != nil {
		//	return mo.Err[any](err)
		//}
		//defer env.DeleteLocalRef(obj)
		//cls := o.Class()
		//return mo.Ok(any(cls))
	case struct{}:
		_, err := env.CallStaticVoidMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(struct{}{}))
	case string:
		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			return mo.Ok(any(""))
		}
		defer env.DeleteLocalRef(obj)
		jstring := env.GetStringUTF(obj)
		return mo.Ok(any(jstring))
	case *string:
		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			var p *string = nil
			return mo.Ok(any(p))
		}
		defer env.DeleteLocalRef(obj)
		jstring := string(env.GetStringUTF(obj))
		return mo.Ok(any(&jstring))
	case int:
		n, err := env.CallStaticIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case int64:
		n, err := env.CallStaticLongMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case bool:
		n, err := env.CallStaticBooleanMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	default:
		panic(errors.New("暂不支持"))
	}
}
func ObjectCallMethod[T JniCallRet](o *Object, funcName string, sig string, args ...any) mo.Result[any] {

	env := LocalThreadJavaEnv()
	objCls, err := o.Class().Get()
	if err != nil {
		return mo.Err[any](err)
	}
	methodId, err := objCls.GetMethodID(funcName, sig).Get()
	if err != nil {
		return mo.Err[any](err)
	}
	jArgs := make([]jni.Jvalue, 0)
	for _, arg := range args {
		jval, needDelete := ConvertAnyArgToJValueArg(arg)
		if needDelete {
			defer env.DeleteLocalRef(jni.Jobject(jval))
		}
		jArgs = append(jArgs, jval)
	}

	var inputType T

	switch any(inputType).(type) {
	case *Object:

		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}

		objWrapper := JniObjectWithPtr[Object](obj)
		return mo.Ok(any(objWrapper))
	case *Class:
		panic(errors.New("system error"))
		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
		//if err != nil {
		//	return mo.Err[any](err)
		//}
		//defer env.DeleteLocalRef(obj)
		//cls := o.Class()
		//return mo.Ok(any(cls))
	case struct{}:
		_, err := env.CallVoidMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(struct{}{}))
	case string:
		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			return mo.Ok(any(""))
		}
		defer env.DeleteLocalRef(obj)
		jstring := env.GetStringUTF(obj)
		return mo.Ok(any(jstring))
	case *string:
		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			var p *string = nil
			return mo.Ok(any(p))
		}
		defer env.DeleteLocalRef(obj)
		jstring := string(env.GetStringUTF(obj))
		return mo.Ok(any(&jstring))
	case int:
		n, err := env.CallIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case int64:
		n, err := env.CallLongMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case bool:
		n, err := env.CallBooleanMethodA(o.JniPtr(), methodId, jArgs...).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	default:

		panic(errors.New("暂不支持"))
	}
}
func GetObjectPtrClassName(ptr jni.Jobject) mo.Result[string] {
	env := LocalThreadJavaEnv()
	classCls, err := Use("java.lang.Class", false).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethod, err := classCls.GetMethodID("getName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}

	cls, err := env.GetObjectClass(ptr).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(cls)
	jstringName, err := env.CallObjectMethodA(cls, getNameMethod).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}
func GetClassName(ptr jni.Jclass) mo.Result[string] {
	env := LocalThreadJavaEnv()
	objCls, err := env.GetObjectClass(ptr).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(objCls)
	getNameMethodId, err := env.GetMethodID(objCls, "getName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	jstringName, err := env.CallObjectMethodA(ptr, getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(jstringName)
	return mo.Ok(string(env.GetStringUTF(jstringName)))
}

func ObjectWrapperGetFieldValue[T JniFieldRet](o *Object, Field string) mo.Result[any] {
	env := LocalThreadJavaEnv()
	objCls, err := o.Class().Get()
	if err != nil {
		return mo.Err[any](err)
	}

	var inputType T

	switch any(inputType).(type) {
	case *Object:
		//fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava.lang.Object").Get()
		//if err != nil {
		//	return mo.Err[any](err)
		//}
		//objWrapper := ObjectWrapperWithJniPtr(obj)
		//return mo.Ok(any(objWrapper))
		panic("no sup")
	case *Class:
		panic(errors.New("system error"))
		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
		//if err != nil {
		//	return mo.Err[any](err)
		//}
		//defer env.DeleteLocalRef(obj)
		//cls := o.Class()
		//return mo.Ok(any(cls))
		panic("no sup")
	case string:
		fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava/lang/String;").Get()
		if err != nil {
			return mo.Err[any](err)
		}
		obj, err := env.GetObjectField(o.JniPtr(), fieldId).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		defer env.DeleteLocalRef(obj)
		jstring := env.GetStringUTF(obj)
		return mo.Ok(any(jstring))
	case *string:
		fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava/lang/String;").Get()
		if err != nil {
			return mo.Err[any](err)
		}
		obj, err := env.GetObjectField(o.JniPtr(), fieldId).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		if obj == 0 {
			var p *string = nil
			return mo.Ok(any(p))
		}
		defer env.DeleteLocalRef(obj)
		jstring := string(env.GetStringUTF(obj))
		return mo.Ok(any(&jstring))
	case int:
		fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava/lang/String;").Get()
		if err != nil {
			return mo.Err[any](err)
		}
		n, err := env.GetIntField(o.JniPtr(), fieldId).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case int64:
		fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava/lang/String;").Get()
		if err != nil {
			return mo.Err[any](err)
		}
		n, err := env.GetLongField(o.JniPtr(), fieldId).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	case bool:
		fieldId, err := env.GetFieldID(objCls.JniPtr(), Field, "Ljava/lang/String;").Get()
		if err != nil {
			return mo.Err[any](err)
		}
		n, err := env.GetBooleanField(o.JniPtr(), fieldId).Get()
		if err != nil {
			return mo.Err[any](err)
		}
		return mo.Ok(any(n))
	default:
		panic(errors.New("暂不支持"))
	}
}

//func ObjectWrapperCall[T JniCallRet](o *Object, funcName string, args ...any) mo.Result[any] {
//	env := LocalThreadJavaEnv()
//	objCls, err := o.Class().Get()
//	if err != nil {
//		return mo.Err[any](err)
//	}
//	method, err := objCls.EnumMatchMethod(funcName, SumGoArgsType(args...)...).Get()
//	if err != nil {
//		return mo.Err[any](err)
//	}
//	defer method.Free()
//
//	methodId, err := env.FromReflectedMethod(method.JniPtr()).Get()
//	if err != nil {
//		return mo.Err[any](err)
//	}
//	jArgs := make([]jni.Jvalue, 0)
//	for _, arg := range args {
//		jval, needDelete := ConvertAnyArgToJValueArg(arg)
//		if needDelete {
//			defer env.DeleteLocalRef(jni.Jobject(jval))
//		}
//		jArgs = append(jArgs, jval)
//	}
//	var inputType T
//
//	switch any(inputType).(type) {
//	case *Object:
//		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		objWrapper := ObjectWrapperWithJniPtr(obj)
//		return mo.Ok(any(objWrapper))
//	case *Class:
//		panic(errors.New("system error"))
//		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
//		//if err != nil {
//		//	return mo.Err[any](err)
//		//}
//		//defer env.DeleteLocalRef(obj)
//		//cls := o.Class()
//		//return mo.Ok(any(cls))
//	case struct{}:
//		_, err := env.CallVoidMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(struct{}{}))
//	case string:
//		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		if obj == 0 {
//			return mo.Ok(any(""))
//		}
//		defer env.DeleteLocalRef(obj)
//		jstring := env.GetStringUTF(obj)
//		return mo.Ok(any(jstring))
//	case *string:
//		obj, err := env.CallObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		if obj == 0 {
//			var p *string = nil
//			return mo.Ok(any(p))
//		}
//		defer env.DeleteLocalRef(obj)
//		jstring := string(env.GetStringUTF(obj))
//		return mo.Ok(any(&jstring))
//	case int:
//		n, err := env.CallIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	case int64:
//		n, err := env.CallLongMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	case bool:
//		n, err := env.CallBooleanMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	default:
//		panic(errors.New("暂不支持"))
//	}
//}
//func ClassWrapperStaticCall[T JniCallRet](o *Class, funcName string, args ...any) mo.Result[any] {
//	env := LocalThreadJavaEnv()
//	method, err := o.EnumMatchStaticMethod(funcName, SumGoArgsType(args...)...).Get()
//	if err != nil {
//		return mo.Err[any](err)
//	}
//	defer method.Free()
//
//	methodId, err := env.FromReflectedMethod(method.JniPtr()).Get()
//	if err != nil {
//		return mo.Err[any](err)
//	}
//	jArgs := make([]jni.Jvalue, 0)
//	for _, arg := range args {
//		jval, needDelete := ConvertAnyArgToJValueArg(arg)
//		if needDelete {
//			defer env.DeleteLocalRef(jni.Jobject(jval))
//		}
//		jArgs = append(jArgs, jval)
//	}
//	var inputType T
//
//	switch any(inputType).(type) {
//	case *Object:
//		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		objWrapper := ObjectWrapperWithJniPtr(obj)
//		return mo.Ok(any(objWrapper))
//	case *Class:
//		panic(errors.New("system error"))
//		//obj, err := o.CallObjectMethodA(methodId, jArgs...).Get()
//		//if err != nil {
//		//	return mo.Err[any](err)
//		//}
//		//defer env.DeleteLocalRef(obj)
//		//cls := o.Class()
//		//return mo.Ok(any(cls))
//	case struct{}:
//		_, err := env.CallStaticVoidMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(struct{}{}))
//	case string:
//		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		if obj == 0 {
//			return mo.Ok(any(""))
//		}
//		defer env.DeleteLocalRef(obj)
//		jstring := env.GetStringUTF(obj)
//		return mo.Ok(any(jstring))
//	case *string:
//		obj, err := env.CallStaticObjectMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		if obj == 0 {
//			var p *string = nil
//			return mo.Ok(any(p))
//		}
//		defer env.DeleteLocalRef(obj)
//		jstring := string(env.GetStringUTF(obj))
//		return mo.Ok(any(&jstring))
//	case int:
//		n, err := env.CallStaticIntMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	case int64:
//		n, err := env.CallStaticLongMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	case bool:
//		n, err := env.CallStaticBooleanMethodA(o.JniPtr(), methodId, jArgs...).Get()
//		if err != nil {
//			return mo.Err[any](err)
//		}
//		return mo.Ok(any(n))
//	default:
//		panic(errors.New("暂不支持"))
//	}
//}

func GetPrimitiveClass(className string) mo.Result[jni.Jclass] {
	env := LocalThreadJavaEnv()
	classClass, err := env.FindClass("java.lang.Class").Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	defer env.DeleteLocalRef(classClass)
	methodId, err := env.GetStaticMethodID(classClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;").Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	ss := env.NewString(className)
	defer env.DeleteLocalRef(ss)
	cls, err := env.CallStaticObjectMethodA(classClass, methodId, jni.Jvalue(ss)).Get()
	if err != nil {
		return mo.Errf[jni.Jclass]("find class error:%v", err)
	}
	return mo.Ok(cls)
}

func As[T any](v IJniObject) *T {
	return JniObjectWithPtr[T](v.JniPtr())
}
