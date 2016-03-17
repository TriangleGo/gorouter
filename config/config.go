package config

import (
	"strconv"
) 


var _globalConfig map[string]string


func InitConfig(){
	_globalConfig = make(map[string]string)
}

func SetConfig(key,value string) {
	_globalConfig[key] = value
}

func GetConfig(key string) string {
	return _globalConfig[key]
}

func GetConfigToInt(key string) int {
	n,_ := strconv.Atoi(_globalConfig[key])
	return n
}
