package util

import (
	"log"
	"runtime"
)

func TraceCrashStack() {
	if x := recover(); x != nil {
		log.Printf("====ERROR:%v====\n", x)
		for i := 0; i < 10; i++ {
			_fn, _file, _line, ok := runtime.Caller(i)
			if ok {
				log.Printf("[%v][%s:%v]\n", runtime.FuncForPC(_fn).Name(), _file, _line)
			}
		}

	}
}
