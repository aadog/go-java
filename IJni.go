package java

import "reflect"

type IJniObject interface {
	SetJniPtr(ptr uintptr)
	JniPtr() uintptr
	DeleteGlobalRef()
	DeleteLocalRef()
}

//	type IJniObject interface {
//		IJni
//		ClassName() mo.Result[string]
//	}
type JniObjectWrapperType interface {
	Object | Class
}

func JniObjectWithPtr[T any](ptr uintptr) *T {
	var t T
	switch tp := any(&t).(type) {
	case *Object:
		tp.SetJniPtr(ptr)
		return &t
	case *Class:
		tp.SetJniPtr(ptr)
		return &t
	default:
		vl := reflect.ValueOf(&t)
		vltp := vl.Type().Elem()

		_, isObj := vltp.FieldByName("Object")
		_, isCls := vltp.FieldByName("Class")
		if isObj {
			vl.Elem().FieldByName("Object").Set(reflect.ValueOf(&Object{ptr: ptr}))
			return &t
		} else if isCls {
			vl.Elem().FieldByName("Class").Set(reflect.ValueOf(&Class{ptr: ptr}))
			return &t
		} else {
			panic("JniObjectWrapperType error")
		}
	}
}
func JniObjectWithPtrAndNewGlobalRef[T JniObjectWrapperType](ptr uintptr) *T {
	env := LocalThreadJavaEnv()
	var t T
	switch tp := any(&t).(type) {
	case *Object:
		tp.SetJniPtr(env.NewGlobalRef(ptr))
		return &t
	case *Class:
		tp.SetJniPtr(env.NewGlobalRef(ptr))
		return &t
	default:
		panic("JniObjectWrapperType error")
	}
}
func JniObjectWithPtrAndNewLocalRef[T JniObjectWrapperType](ptr uintptr) *T {
	env := LocalThreadJavaEnv()
	var t T
	switch tp := any(&t).(type) {
	case *Object:
		tp.SetJniPtr(env.NewLocalRef(ptr))
		return &t
	case *Class:
		tp.SetJniPtr(env.NewLocalRef(ptr))
		return &t
	default:
		panic("JniObjectWrapperType error")
	}
}
