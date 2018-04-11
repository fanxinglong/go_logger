package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level    int    //日志级别
	LogPath  string //日志文件路径
	LogName  string //日志文件名称
	file     *os.File
	warnFile *os.File
}

//构造函数
func NewFileLogger(level int, LogPath, LogName string) LogInterface {
	logger := &FileLogger{
		level:   level,
		LogPath: LogPath,
		LogName: LogName,
	}
	//调用初始化方法
	logger.init()
	return logger
}

//初始化
func (f *FileLogger) init() {
	//打印debug 正常信息的文件
	filename := fmt.Sprintf("%s/%s.log", f.LogPath, f.LogName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file%s is faild,error:%v", filename, err))
	}
	f.file = file
	//打印error 和fatal信息的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.LogPath, f.LogName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file%s is faild,error:%v", filename, err))
	}
	f.warnFile = file

}

//实现接口规范的方法
func (f *FileLogger) SetLevel(level int) {
	//级别范围判断
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) writeLog(file *os.File, level int, format string, args ...interface{}) {
	if f.level > level {
		return
	}
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

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.writeLog(f.file, LogLevelDebug, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	f.writeLog(f.file, LogLevelTrace, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.writeLog(f.file, LogLevelInfo, format, args...)
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.writeLog(f.warnFile, LogLevelWarn, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.writeLog(f.warnFile, LogLevelError, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.writeLog(f.warnFile, LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
