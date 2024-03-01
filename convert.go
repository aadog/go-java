package java

import (
	"errors"
	"fmt"
	"github.com/aadog/go-jni"
	"github.com/samber/lo"
	"reflect"
	//"strings"
)

//func DeclaredConstructorsToString(items []*JavaLangReflectConstructorObjectWrapper) string {
//	sMethods := strings.Builder{}
//	for _, method := range items {
//		sMethods.WriteString("\n")
//		sMethods.WriteString(method.ToString().OrElse(""))
//		sMethods.WriteString("\n")
//	}
//	return sMethods.String()
//}
//func DeclaredMethodsToString(items []*JavaLangReflectMethodObjectWrapper) string {
//	sMethods := strings.Builder{}
//	for _, method := range items {
//		sMethods.WriteString("\n")
//		sMethods.WriteString(method.ToString().OrElse(""))
//		sMethods.WriteString("\n")
//	}
//	return sMethods.String()
//}

func ConvertAnyArgToJValueArg(arg any) (jni.Jvalue, bool) {
	env := LocalThreadJavaEnv()
	var jValArg jni.Jvalue
	vl := reflect.ValueOf(arg)
	tp := vl.Type()
	//if tp.Kind() == reflect.Slice || tp.Kind() == reflect.Array {
	//	return fmt.Sprintf("%s%s", ConvertGoBaseTypeToJavaBaseType(tp.Elem()), "[]")
	//}
	if tp.Kind() == reflect.String {

		return jni.Jvalue(env.NewString(vl.String())), true
	}
	if tp.Kind() == reflect.Int {
		return jni.Jvalue(vl.Int()), false
	}
	if tp.Kind() == reflect.Int8 {
		return jni.Jvalue(int8(vl.Int())), false
	}
	if tp.Kind() == reflect.Uint8 {
		return jni.Jvalue(uint8(vl.Uint())), false
	}
	if tp.Kind() == reflect.Int16 {
		return jni.Jvalue(int16(vl.Int())), false
	}
	if tp.Kind() == reflect.Uint16 {
		return jni.Jvalue(uint16(vl.Uint())), false
	}
	if tp.Kind() == reflect.Int32 {
		return jni.Jvalue(int32(vl.Int())), false
	}
	if tp.Kind() == reflect.Uint32 {
		return jni.Jvalue(uint32(vl.Uint())), false
	}
	if tp.Kind() == reflect.Int64 {
		return jni.Jvalue(vl.Int()), false
	}
	if tp.Kind() == reflect.Uint64 {
		return jni.Jvalue(vl.Uint()), false
	}
	if tp.Kind() == reflect.Float32 {
		return jni.Jvalue(float32(vl.Float())), false
	}
	if tp.Kind() == reflect.Float64 {
		return jni.Jvalue(float64(vl.Float())), false
	}
	if tp.Kind() == reflect.Bool {
		return jni.Jvalue(lo.If(vl.Bool(), 1).Else(0)), false
	}
	if tp.Kind() == reflect.Pointer {
		obj, ok := vl.Interface().(IJniObject)
		if ok {
			return jni.Jvalue(obj.JniPtr()), false
		} else {
			panic(errors.New(fmt.Sprintf("error type:%v", vl.Interface())))
		}
	}
	panic(errors.New("convertAnyArgToJValueArg: invalid argument type"))
	return jValArg, false
}
func ConvertGoBaseTypeToJavaBaseType(tp reflect.Type) string {
	if tp.Kind() == reflect.Slice || tp.Kind() == reflect.Array {
		return fmt.Sprintf("%s%s", ConvertGoBaseTypeToJavaBaseType(tp.Elem()), "[]")
	}
	if tp.Kind() == reflect.String {
		return "java.lang.String"
	}
	if tp.Kind() == reflect.Int {
		return "int"
	}
	if tp.Kind() == reflect.Int8 {
		return "byte"
	}
	if tp.Kind() == reflect.Uint8 {
		return "byte"
	}
	if tp.Kind() == reflect.Int16 {
		return "short"
	}
	if tp.Kind() == reflect.Uint16 {
		return "short"
	}
	if tp.Kind() == reflect.Int32 {
		return "int"
	}
	if tp.Kind() == reflect.Uint32 {
		return "int"
	}
	if tp.Kind() == reflect.Int64 {
		return "long"
	}
	if tp.Kind() == reflect.Uint64 {
		return "long"
	}
	if tp.Kind() == reflect.Float32 {
		return "float"
	}
	if tp.Kind() == reflect.Float64 {
		return "double"
	}
	if tp.Kind() == reflect.Bool {
		return "boolean"
	}
	panic(errors.New("convert error"))
	return tp.Name()
}
