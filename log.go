package java

/*
#include <android/log.h>
*/
import "C"
import (
	"fmt"
)

func LogError(tag, format string, v ...interface{}) {
	ctag := (*C.char)(CString(tag))
	cstr := (*C.char)(CString(fmt.Sprintf(format, v...)))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}

func LogInfo(tag, format string, v ...interface{}) {
	ctag := (*C.char)(CString(tag))
	cstr := (*C.char)(CString(fmt.Sprintf(format, v...)))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
}
