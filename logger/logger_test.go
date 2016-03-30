package logger

import (

	"testing"

)



func Test_Logger(t *testing.T) {
	GetLogger().SetLogLevel(5,5)
	GetLogger().SetOutputDir(`c:\`)
	GetLogger().Init("testlog")
	xx := make([]byte,9999)
	Fn := func() {
		for i:=0;i<9999;i++ {
		Test("padding some data %d \n",i)
		Debug("padding some data %d \n",i)
		Info("padding some data %d \n",i)
		Error("padding some data %d \n",i)
		Critial("padding some data %d \n",i)
		Info("data :%v\n",xx)
	}
	
	}
	for i:=0;i<9999;i++ {
		go Fn()
	}

	
}
