package java

import "github.com/samber/mo"

type JavaUtilHashMapObject struct {
	*Object
}

func (j *JavaUtilHashMapObject) Size() mo.Result[int] {
	return j.CallIntA("size", "()I")
}
func (j *JavaUtilHashMapObject) IsEmpty() mo.Result[bool] {
	return j.CallBooleanA("isEmpty", "()Z")
}
func (j *JavaUtilHashMapObject) Clear() mo.Result[bool] {
	return j.CallBooleanA("clear", "()V")
}
func (j *JavaUtilHashMapObject) Put(key IJniObject, v IJniObject) mo.Result[struct{}] {
	o, err := j.CallObjectA("put", "(Ljava/lang/Object;Ljava/lang/Object;)Ljava/lang/Object;", key, v).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	o.DeleteLocalRef()
	return mo.Ok(struct{}{})
}
func (j *JavaUtilHashMapObject) Remove(key IJniObject) mo.Result[struct{}] {
	o, err := j.CallObjectA("remove", "(Ljava/lang/Object;)Ljava/lang/Object;", key).Get()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	o.DeleteLocalRef()
	return mo.Ok(struct{}{})
}
func (j *JavaUtilHashMapObject) Get(key IJniObject) mo.Result[*Object] {
	o, err := j.CallObjectA("get", "(Ljava/lang/Object;)Ljava/lang/Object;", key).Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	return mo.Ok(o)
}
func (j *JavaUtilHashMapObject) ContainsKey(key IJniObject) mo.Result[bool] {
	o, err := j.CallBooleanA("containsKey", "(Ljava/lang/Object;)Z", key).Get()
	if err != nil {
		return mo.Err[bool](err)
	}
	return mo.Ok(o)
}

func (j *JavaUtilHashMapObject) KeySet(key IJniObject) mo.Result[*Object] {
	o, err := j.CallObjectA("keySet", "()Ljava/lang/Object;", key).Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	return mo.Ok(o)
}
