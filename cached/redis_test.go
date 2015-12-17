package cached

import (
	"time"
	"testing"
	"fmt"
	_"github.com/garyburd/redigo/redis"
	"gorouter/logger"
)



func Test_Redis(t *testing.T) {
	logger.GetLogger().SetLogLevel(5,5)
	logger.GetLogger().SetOutputDir(`c:\`)
	logger.GetLogger().Init("testlog")
	fmt.Printf("redis test running \n")
	for {
		v:= GetCached().Do("get","abc")
		fmt.Printf("testing v =%v err %v \n",v)
		time.Sleep(1 * time.Second)
	}

}


