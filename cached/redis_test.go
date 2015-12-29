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
	
	rd := NewRedis()
	rd.Init()
	ret ,_ := rd.Pool().Get().Do("GET","abc")
	
	
	fmt.Printf("redis test running %v\n",ret)
	for {
		v:= GetCached().Do("get","abc")
		fmt.Printf("testing v =%v  \n",v)
		time.Sleep(1 * time.Second)
	}

}


