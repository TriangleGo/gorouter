package config

import (
	"strconv"
) 


var _globalConfig map[string]string


func InitConfig(){
	_globalConfig = make(map[string]string)
	_globalConfig["redis_host"] = "127.0.0.1:6379"

	_globalConfig["mysql_host"] = "127.0.0.1:51816"
	_globalConfig["mysql_user"] = "oa_local"
	_globalConfig["mysql_pass"] = `f*(&Dssdsa)s`
	_globalConfig["mysql_dbname"] = "db_oa_enterprise"
	
	
	
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