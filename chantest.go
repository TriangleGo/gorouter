package main



import (
	_"strings"
	_"time"
	_"os"
	"fmt"
	_"runtime"
	_"path/filepath"
	_"gorouter/logger"
	"github.com/garyburd/redigo/redis"
	"gorouter/lib/goconfig"
	"encoding/json"
)


type Test struct  {
	Name string 
	Age int 
	Userid int
}

func main () {
	fmt.Printf("logger test \n")
	c, err_config := goconfig.ReadConfigFile("config.conf"); if err_config != nil {
        fmt.Printf("config file %s \n",err_config)
    	}
		
	fmt.Printf("config file %v \n ",c)
	logLv,_ := c.GetInt64("log", "file_log_level");	
	fmt.Printf("config file logLv %v \n",logLv)

	pool := redis.NewPool(func()(redis.Conn, error){
			return redis.Dial("tcp", "127.0.0.1:6379")},10)
	pool.MaxActive = 10
	
	t := Test{Name:"John",Age:16,Userid:10086}
	tJson, err := json.Marshal(t)

	
	conn := pool.Get()
	i,err := conn.Do(`hset`,"domain","key",tJson)
	fmt.Printf("i %v error %v \n",i,err)
	
	i,err = conn.Do(`hget`,"domain","key")
	
	
	
	fmt.Printf("ret string %s \n",i)
	nt := Test{}
	err = json.Unmarshal(i.([]byte),&nt)
	fmt.Printf("get %v %v\n", nt,err)
	conn.Close()
}




