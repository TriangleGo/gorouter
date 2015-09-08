package logger

import (
	
)

var _logger *Logger

func GetLogger() *Logger {
	if _logger == nil {
		_logger = NewLogger()
	}
	return _logger
}

func Test(format string,v ...interface{}) {
	GetLogger().Log(5,format,v...)
}

func Debug(format string,v ...interface{}) {
	GetLogger().Log(4,format,v...)
}

func Info(format string,v ...interface{}) {
	GetLogger().Log(3,format,v...)
}

func Error(format string,v ...interface{}) {
	GetLogger().Log(2,format,v...)
}

func Critial(format string,v ...interface{}) {
	GetLogger().Log(1,format,v...)
}

