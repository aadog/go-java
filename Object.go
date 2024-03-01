package java

import "C"
import (
	"github.com/samber/mo"
)

type Object struct {
	ptr uintptr
}

func (o *Object) DeleteGlobalRef() {
	env := LocalThreadJavaEnv()
	env.DeleteGlobalRef(o.ptr)
}
func (o *Object) DeleteLocalRef() {
	env := LocalThreadJavaEnv()
	env.DeleteLocalRef(o.ptr)
}
func (o *Object) JniPtr() uintptr {
	return o.ptr
}
func (o *Object) SetJniPtr(ptr uintptr) {
	o.ptr = ptr
}
func (o *Object) ClassName() mo.Result[string] {
	return GetObjectPtrClassName(o.JniPtr())
}
func (o *Object) Class() mo.Result[*Class] {
	clsName, err := o.ClassName().Get()
	if err != nil {
		return mo.Err[*Class](err)
	}
	cls, err := Use(clsName, false).Get()
	if err != nil {
		if clsName != "" {
			env := LocalThreadJavaEnv()
			StoreClass(clsName, env.GetObjectClass(o.JniPtr()).MustGet())
			return Use(clsName, false)
		} else {
			return mo.Err[*Class](err)
		}
	}
	return mo.Ok(cls)
}
func (o *Object) CallObjectA(funcName string, sig string, args ...any) mo.Result[*Object] {
	ret, err := ObjectCallMethod[*Object](o, funcName, sig, args...).Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	return mo.Ok(ret.(*Object))
}
func (o *Object) CallIntA(funcName string, sig string, args ...any) mo.Result[int] {
	ret, err := ObjectCallMethod[int](o, funcName, sig, args...).Get()
	if err != nil {
		return mo.Err[int](err)
	}
	return mo.Ok(ret.(int))
}
func (o *Object) CallBooleanA(funcName string, sig string, args ...any) mo.Result[bool] {
	ret, err := ObjectCallMethod[bool](o, funcName, sig, args...).Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(ret.(bool))
}
func (o *Object) CallPStringA(funcName string, sig string, args ...any) mo.Result[*string] {
	ret, err := ObjectCallMethod[*string](o, funcName, sig, args...).Get()
	if err != nil {
		return mo.Err[*string](err)
	}
	return mo.Ok(ret.(*string))
}

func (o *Object) GetStringFieldValue(funcName string) mo.Result[string] {
	ret, err := ObjectWrapperGetFieldValue[string](o, funcName).Get()
	if err != nil {
		return mo.Err[string](err)
	}
	return mo.Ok(ret.(string))
}

//func (o *Object) CallPStringA(funcName string, args ...any) mo.Result[*string] {
//	ret, err := ObjectWrapperCall[*string](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[*string](err)
//	}
//	return mo.Ok(ret.(*string))
//}
//func (o *Object) CallStringA(funcName string, args ...any) mo.Result[string] {
//	ret, err := ObjectWrapperCall[string](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[string](err)
//	}
//	return mo.Ok(ret.(string))
//}

//func (o *Object) CallVoidA(funcName string, args ...any) mo.Result[struct{}] {
//	ret, err := ObjectWrapperCall[struct{}](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[struct{}](err)
//	}
//	return mo.Ok(ret.(struct{}))
//}
//func (o *Object) CallIntA(funcName string, args ...any) mo.Result[int] {
//	ret, err := ObjectWrapperCall[int](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[int](err)
//	}
//	return mo.Ok(ret.(int))
//}
//func (o *Object) CallLongA(funcName string, args ...any) mo.Result[int64] {
//	ret, err := ObjectWrapperCall[int64](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[int64](err)
//	}
//	return mo.Ok(ret.(int64))
//}
//func (o *Object) CallBoolA(funcName string, args ...any) mo.Result[bool] {
//	ret, err := ObjectWrapperCall[bool](o, funcName, args...).Get()
//	if err != nil {
//		return mo.Err[bool](err)
//	}
//	return mo.Ok(ret.(bool))
//}

func (o *Object) ToString() mo.Result[string] {
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
func Cast(ptr uintptr, class string) *Class {
	return &Class{ptr: ptr}
}
