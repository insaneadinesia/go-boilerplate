package helper

import (
	"runtime"
	"strings"
)

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		name := details.Name()
		index := strings.LastIndex(name, "/")
		if index >= 0 {
			return name[index+1:]
		}

		return name
	}

	return ""
}
