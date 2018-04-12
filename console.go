package logger

import (
	"os"
)

type ConsoleLogger struct {
	level int //日志级别
}

//构造函数
func NewConsoleLogger(level int) LogInterface {
	logger := &ConsoleLogger{
		level: level,
	}
	return logger
}

//实现接口规范的方法
func (c *ConsoleLogger) SetLevel(level int) {
	//级别范围判断
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	//writeLog(f.file, LogLevelDebug, format, args...)
	if c.level > LogLevelDebug {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)

}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	writeLog(os.Stdout, LogLevelTrace, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	writeLog(os.Stdout, LogLevelInfo, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	writeLog(os.Stdout, LogLevelWarn, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	writeLog(os.Stdout, LogLevelError, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	writeLog(os.Stdout, LogLevelFatal, format, args...)
}

func (f *ConsoleLogger) Close() {

}
