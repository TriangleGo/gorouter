package main



import (
	_"strings"
	"time"
	_"os"
	"fmt"
	_"runtime"
	_"path/filepath"
	"./lib/goconfig"
	"gorouter/logger"
)


func main () {
	fmt.Printf("logger test \n")
	c, err_config := goconfig.ReadConfigFile("config.conf"); if err_config != nil {
        fmt.Printf("config file %s \n",err_config)
    	}
		
	fmt.Printf("config file %v \n ",c)
	logLv,_ := c.GetInt64("log", "file_log_level");	
	fmt.Printf("config file logLv %v \n",logLv)

	log := logger.NewLogger()
	log.SetOutputDir(`c:\`)
	log.SetLogLevel(5,5)
	log.Init("log")
	for i:=0;i<99999;i++ {
		log.Log(5,"i am the logger %v %d\n","hahahaha",i)	
		time.Sleep(time.Millisecond * 1000)
	}
	log.Close()
	
	
	fmt.Printf("asdasd %v \n","123")
	
}




