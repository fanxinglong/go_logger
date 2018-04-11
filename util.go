package logger

import (
	"runtime"
)

// 获取调用文件名称 行号  函数名称
func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}
