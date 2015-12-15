package jsonconfig

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
)

type JsonConfig struct {
	json *simplejson.Json
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{}
}

func (this *JsonConfig)LoadConfigFile(filepath string) {
	byteConf ,err := ioutil.ReadFile(filepath)
	if err != nil {
		//log.Fatal("error loading db connect setting\n")
		fmt.Printf("Can't load config file reason:%v \n",err)
		return
	}
	
	this.json ,err = simplejson.NewJson(byteConf)
	
	if err != nil {
		fmt.Printf("Can't parse config file reason:%v \n",err)
		return
	}
}

func (this *JsonConfig) GetValue(key string) string{
	return this.json.Get(key).MustString()
}

func (this *JsonConfig) GetValueEx(section,key string) string{
	return this.json.Get(section).Get(key).MustString()
}

func (this *JsonConfig) GetInt(key string) int {
	return this.json.Get(key).MustInt()
}

func (this *JsonConfig) GetIntEx(section,key string) int{
	return this.json.Get(section).Get(key).MustInt()
}