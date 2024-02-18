package java

import "github.com/samber/mo"

type IJni interface {
	JniPtr() uintptr
}

type IJniObject interface {
	IJni
	ClassName() mo.Result[string]
}
