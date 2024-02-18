package java

import "C"
import (
	"github.com/samber/mo"
)

type ObjectWrapper struct {
	ptr      uintptr
	isGlobal bool
}

// func (c *ClassWrapper) New(args ...any) mo.Result[*ObjectWrapper] {
//
//		constructor, err := c.MatchConstructor(SumGoArgsType(args...)...).Get()
//		if err != nil {
//			return mo.Err[*ObjectWrapper](err)
//		}
//		defer constructor.Free()
//		env := LocalThreadJavaEnv()
//		constructorMethodId, err := env.FromReflectedMethod(constructor.JniPtr()).Get()
//		if err != nil {
//			return mo.Err[*ObjectWrapper](err)
//		}
//
//		a := env.NewString("http://www.baidu.com")
//		defer env.DeleteLocalRef(a)
//		obj, err := env.NewObjectA(c.JniPtr(), constructorMethodId, jni.Jvalue(a)).Get()
//		if err != nil {
//			return mo.Err[*ObjectWrapper](err)
//		}
//		LogError("Go", "obj:%v", obj)
//
//		objWrapper := ObjectWrapperWithJniPtr(env.NewGlobalRef(obj))
//		runtime.SetFinalizer(objWrapper, func(obj *ObjectWrapper) {
//			obj.Free()
//		})
//		return mo.Ok[*ObjectWrapper](objWrapper)
//	}

func (o *ObjectWrapper) GetStringFieldValue(funcName string) mo.Result[string] {
	ret, err := ObjectWrapperGetFieldValue[string](o, funcName).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(ret.(string))
}
func (o *ObjectWrapper) CallPStringA(funcName string, args ...any) mo.Result[*string] {
	ret, err := ObjectWrapperCall[*string](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[*string](err)
	}
	return mo.Ok(ret.(*string))
}
func (o *ObjectWrapper) CallStringA(funcName string, args ...any) mo.Result[string] {
	ret, err := ObjectWrapperCall[string](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(ret.(string))
}
func (o *ObjectWrapper) CallObjectA(funcName string, args ...any) mo.Result[*ObjectWrapper] {
	ret, err := ObjectWrapperCall[*ObjectWrapper](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[*ObjectWrapper](err)
	}
	return mo.Ok(ret.(*ObjectWrapper))
}
func (o *ObjectWrapper) CallVoidA(funcName string, args ...any) mo.Result[struct{}] {
	ret, err := ObjectWrapperCall[struct{}](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(ret.(struct{}))
}
func (o *ObjectWrapper) CallIntA(funcName string, args ...any) mo.Result[int] {
	ret, err := ObjectWrapperCall[int](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(ret.(int))
}
func (o *ObjectWrapper) CallLongA(funcName string, args ...any) mo.Result[int64] {
	ret, err := ObjectWrapperCall[int64](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[int64](err)
	}
	return mo.Ok(ret.(int64))
}
func (o *ObjectWrapper) CallBoolA(funcName string, args ...any) mo.Result[bool] {
	ret, err := ObjectWrapperCall[bool](o, funcName, args...).Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(ret.(bool))
}

func (o *ObjectWrapper) ToString() mo.Result[string] {
	env := LocalThreadJavaEnv()
	cls, err := o.Class().Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethodId, err := cls.GetMethodID("toString", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	ret, err := env.CallObjectMethodA(o.JniPtr(), getNameMethodId).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	defer env.DeleteLocalRef(ret)
	return mo.Ok(env.GetStringUTF(ret))
}

func (o *ObjectWrapper) ClassName() mo.Result[string] {
	env := LocalThreadJavaEnv()
	classCls, err := Use("java.lang.Class").Get()
	if err != nil {
		return mo.Err[string](err)
	}
	getNameMethod, err := classCls.GetMethodID("getName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[string](err)
	}

	cls, err := env.GetObjectClass(o.JniPtr()).Get()
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
func (o *ObjectWrapper) Class() mo.Result[*ClassWrapper] {
	clsName, err := o.ClassName().Get()
	if err != nil {
		return mo.Err[*ClassWrapper](err)
	}
	cls, err := Use(clsName).Get()
	if err != nil {
		if clsName != "" {
			env := LocalThreadJavaEnv()
			return ForeUse(clsName, env.GetObjectClass(o.JniPtr()).MustGet())
		} else {
			return mo.Err[*ClassWrapper](err)
		}
	}
	return mo.Ok(cls)
}
func (o *ObjectWrapper) Free() {
	env := LocalThreadJavaEnv()
	if o.isGlobal {
		env.DeleteGlobalRef(o.ptr)
	} else {
		env.DeleteLocalRef(o.ptr)
	}
}
func (o *ObjectWrapper) JniPtr() uintptr {
	return o.ptr
}
func Cast(ptr uintptr, class string) *ClassWrapper {
	return &ClassWrapper{ptr: ptr}
}

func ObjectWrapperWithJniPtr(ptr uintptr) *ObjectWrapper {
	return &ObjectWrapper{ptr: ptr, isGlobal: false}
}

func ObjectWrapperWithGlobalJniPtr(ptr uintptr) *ObjectWrapper {
	return &ObjectWrapper{ptr: ptr, isGlobal: true}
}
