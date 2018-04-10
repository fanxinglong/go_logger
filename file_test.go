package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/Users/xinglongfan/Desktop", "test")
	logger.Debug("user id[%d] is from china", 23132)
	logger.Warn("user id[%d] is from china", 54534)
	logger.Fatal("user id[%d] is from china", 78889)
	logger.Close()

}
