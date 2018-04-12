package logger

import (
	"testing"
)

//文件日志打印测试
func TestFileLogger(t *testing.T) {
	//对于客户端调用者来讲 我只需要构造一个日志对象
	//指定日志打印级别，日志打印的展现方式，支持格式化打印
	logger := NewFileLogger(LogLevelDebug, "/Users/xinglongfan/Desktop", "test")
	logger.Debug("user id[%d] is from china", 23132)
	logger.Warn("user id[%d] is from china", 54534)
	logger.Fatal("user id[%d] is from china", 78889)
	logger.Close()
}

//终端日志打印测试
func TestConsoleLogger(t *testing.T) {
	//对于客户端调用者来讲 我只需要构造一个日志对象
	//指定日志打印级别，日志打印的展现方式，支持格式化打印
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is from china", 23132)
	logger.Warn("user id[%d] is from china", 54534)
	logger.Fatal("user id[%d] is from china", 78889)
	logger.Close()
}
