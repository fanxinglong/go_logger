package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

// 获取调用文件名称 行号  函数名称
func getLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func writeLog(file *os.File, level int, format string, args ...interface{}) {
	//str := fmt.Sprintf(format, args...)
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("2006-01-02 15:04:05")
	levelStr := getLevelToText(level)
	fileName, funcName, lineNo := getLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s,%s,(%s:%s:%d),%s\n", timeNow, levelStr, fileName, funcName, lineNo, msg)
}
