package cached

import (
	"time"
	
	"github.com/garyburd/redigo/redis"
	"github.com/TriangleGo/gorouter/config"
	"github.com/TriangleGo/gorouter/logger"
)


var _redis *Redis

func GetCached() *Redis {
	if _redis == nil {
		_redis = NewRedis()
		_redis.Init()
	}
	return _redis
}


type Redis struct {
	pool  *redis.Pool
	maxIdle int
	maxActive int
}


func NewRedis() *Redis{
	return &Redis{}
}


func (this *Redis) Init() {
	this.pool = redis.NewPool(func()(redis.Conn, error){
			return redis.Dial("tcp", config.GetConfig("redis_host"))},20)
	this.pool.IdleTimeout = 45 * time.Second
	this.pool.MaxActive = 30
	this.pool.MaxIdle = 30
} 

func (this *Redis) Do(command string, args ...interface{} ) (interface{}) {
	conn := this.pool.Get()
	defer conn.Close()
	reply , err :=conn.Do(command,args ...)
	
	if err != nil {
		logger.Error("Redis command error :%v\n",err)
		return nil
	}
	return reply
} 

func (this *Redis) Keys(pattern string) ([]string) {
	reply := this.Do("keys",pattern)
	retCount := len( reply.([]interface{}) )
	if retCount == 0 {
		return nil
	}
	
	var keysArray []string
	for _,v := range  reply.([]interface{}) {
		keysArray = append(keysArray, string(v.([]byte)) )
	}
	
	return keysArray
}


/*
func (this *Redis) Pool() *redis.Pool {
	return this.pool
}
*/
