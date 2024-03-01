package java

import "github.com/samber/mo"

func GetApplicationContext() mo.Result[*Object] {
	androidAppActivityThreadClass, err := Use("android.app.ActivityThread", false).Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	currentActivityThread, err := androidAppActivityThreadClass.CallStaticObjectA("currentActivityThread", "()Landroid/app/ActivityThread;").Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	defer currentActivityThread.DeleteLocalRef()
	application, err := currentActivityThread.CallObjectA("getApplication", "()Landroid/app/Application;").Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	defer application.DeleteLocalRef()
	applicationContext, err := application.CallObjectA("getApplicationContext", "()Landroid/content/Context;").Get()
	if err != nil {
		return mo.Err[*Object](err)
	}
	return mo.Ok(applicationContext)
}
func GetPackageName(applicationContext *Object) mo.Result[*string] {
	packageName, err := applicationContext.CallPStringA("getPackageName", "()Ljava/lang/String;").Get()
	if err != nil {
		return mo.Err[*string](err)
	}
	return mo.Ok(packageName)
}
