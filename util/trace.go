package util

import (
	"runtime"
	"github.com/TriangleGo/gorouter/logger"
)

func TraceCrashStack() {
	if x := recover(); x != nil {
		logger.Error("====ERROR:%v====\n", x)
		for i := 0; i < 10; i++ {
			_fn, _file, _line, ok := runtime.Caller(i)
			if ok {
				logger.Error("[%v][%s:%v]\n", runtime.FuncForPC(_fn).Name(), _file, _line)
			}
		}
	}
}


func TraceSupervisor( fn func() ) {
	if x := recover(); x != nil {
		logger.Error("====ERROR:%v====\n", x)
		for i := 0; i < 10; i++ {
			_fn, _file, _line, ok := runtime.Caller(i)
			if ok {
				logger.Error("[%v][%s:%v]\n", runtime.FuncForPC(_fn).Name(), _file, _line)
			}
		}
		fn()
		
	}
}
