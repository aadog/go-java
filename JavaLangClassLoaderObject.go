package java

import "github.com/samber/mo"
import "strings"

type JavaLangClassLoaderObject struct {
	*Object
}

func (j *JavaLangClassLoaderObject) LoadClass(clsName string) mo.Result[*Class] {
	clsName = strings.ReplaceAll(clsName, ".", "/")
	o, err := j.CallObjectA("loadClass", "(Ljava/lang/String;)Ljava/lang/Class;", clsName).Get()
	if err != nil {
		return mo.Err[*Class](err)
	}
	StoreClass(clsName, o.JniPtr())
	o.DeleteLocalRef()
	return Use(clsName, false)
}
func (j *JavaLangClassLoaderObject) GetParent() mo.Result[*JavaLangClassLoaderObject] {
	o, err := j.CallObjectA("getParent", "()Ljava/lang/ClassLoader;").Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	return mo.Ok(As[JavaLangClassLoaderObject](o))
}
func (j *JavaLangClassLoaderObject) FindClass(clsName string) mo.Result[*Class] {
	clsName = strings.ReplaceAll(clsName, ".", "/")
	o, err := j.CallObjectA("findClass", "(Ljava/lang/String;)Ljava/lang/Class;", clsName).Get()
	if err != nil {
		return mo.Err[*Class](err)
	}
	StoreClass(clsName, o.JniPtr())
	o.DeleteLocalRef()
	return Use(clsName, false)
}
func (j *JavaLangClassLoaderObject) FindLibrary(libName string) mo.Result[*string] {
	o, err := j.CallPStringA("findLibrary", "(Ljava/lang/String;)Ljava/lang/String;", libName).Get()
	if err != nil {
		return mo.Err[*string](err)
	}
	return mo.Ok(o)
}
