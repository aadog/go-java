package java

import "github.com/samber/mo"

type JavaLangStringObject struct {
	*Object
}

func (j *JavaLangClassLoaderObject) Length() mo.Result[int] {
	return j.CallIntA("length", "()I")
}
