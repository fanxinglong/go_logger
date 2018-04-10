package logger

/*
日志库接口规范
*/
type LogInterface interface {
	SetLevel(level int) //设置日志库级别
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
