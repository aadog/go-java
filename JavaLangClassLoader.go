package java

import (
	"fmt"
	"github.com/samber/mo"
)

type JavaLangClassLoader struct {
	*Class
}

func (j *JavaLangClassLoader) GetSystemClassLoader() mo.Result[*JavaLangClassLoaderObject] {
	o, err := j.CallStaticObjectA("getSystemClassLoader", "()Ljava/lang/ClassLoader;").Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	return mo.Ok(As[JavaLangClassLoaderObject](o))
}

func (j *JavaLangClassLoader) GetSystemResource() mo.Result[*string] {
	return j.CallStaticPStringA("getSystemResource", "(Ljava/lang/String;)Ljava/net/URL;")
}

func (j *JavaLangClassLoader) GetSystemResourceAsStream() mo.Result[*string] {
	return j.CallStaticPStringA("getSystemResourceAsStream", "(Ljava/lang/String;)Ljava/io/InputStream;")
}

func DalvikSystemDexClassLoaderWithDexPath(dexPath string) mo.Result[*JavaLangClassLoaderObject] {
	javaLangClassLoader, err := UseT[JavaLangClassLoader]("java.lang.ClassLoader", false).Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	systemClassLoader, err := javaLangClassLoader.GetSystemClassLoader().Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	defer systemClassLoader.DeleteLocalRef()
	dalvikSystemDexClassLoader, err := Use("dalvik.system.DexClassLoader", false).Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	context, err := GetApplicationContext().Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	defer context.DeleteLocalRef()
	packageName, err := GetPackageName(context).Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	cacheDir := fmt.Sprintf("/data/data/%s/code_cache", packageName)
	newClassLoader, err := dalvikSystemDexClassLoader.New("(Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;Ljava/lang/ClassLoader;)V", dexPath, cacheDir, 0, systemClassLoader).Get()
	if err != nil {
		return mo.Err[*JavaLangClassLoaderObject](err)
	}
	env := LocalThreadJavaEnv()
	if env.ExceptionCheck() {
		panic(env.GetAndClearExceptionMessage())
	}
	return mo.Ok[*JavaLangClassLoaderObject](As[JavaLangClassLoaderObject](newClassLoader))
}
