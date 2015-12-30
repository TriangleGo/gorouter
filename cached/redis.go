package cached

import (
	
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
	this.pool.MaxActive = 0
} 

func (this *Redis) Do(command string, args ...interface{} ) (interface{}) {
	reply , err := this.pool.Get().Do(command,args ...)
	
	logger.Debug("err %v \n",err)
	
	if err != nil {
		logger.Error("Redis command error :%v\n",err)
		return nil
	}
	return reply
} 

/*
func (this *Redis) Pool() *redis.Pool {
	return this.pool
}
*/
