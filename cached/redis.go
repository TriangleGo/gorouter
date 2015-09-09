package cached

import (
	"github.com/garyburd/redigo/redis"
	_"gorouter/lib/goconfig"
	_"encoding/json"
)


var _redis *Redis

func GetCached() *Redis {
	if _redis = nil {
		r := NewRedis()
		r.Init()
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
	this.pool := redis.NewPool(func()(redis.Conn, error){
			return redis.Dial("tcp", "127.0.0.1:6379")},10)
	this.pool.MaxActive = 10
} 