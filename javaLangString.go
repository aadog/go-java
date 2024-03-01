package java

type JavaLangString struct {
	*Class
}

func (j *JavaLangString) NewWithString(s string) *JavaLangStringObject {
	env := LocalThreadJavaEnv()
	js := env.NewString(s)
	return JniObjectWithPtr[JavaLangStringObject](js)
}

func NewJString(s string) *JavaLangStringObject {
	javaLangString := UseT[JavaLangString]("java.lang.String", false).MustGet()
	return javaLangString.NewWithString(s)
}
