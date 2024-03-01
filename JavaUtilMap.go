package java

import "github.com/samber/mo"

type JavaUtilMap struct {
	*Class
}

func (j *JavaUtilMap) New() mo.Result[*JavaUtilMapObject] {
	o, err := j.Class.New("()V").Get()
	if err != nil {
		return mo.Err[*JavaUtilMapObject](err)
	}
	return mo.Ok(As[JavaUtilMapObject](o))
}
