package cached

import (
	"time"

	"github.com/TriangleGo/gorouter/config"
	//"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/redisgo/redis"
)

var _ssdb *Ssdb

func GetSsdb() *Ssdb {
	if _ssdb == nil {
		_ssdb = NewSsdb()
		_ssdb.Init()
	}
	return _ssdb
}

type Ssdb struct {
	pool      *redis.Pool
	maxIdle   int
	maxActive int
}

func NewSsdb() *Ssdb {
	return &Ssdb{}
}

func (this *Ssdb) Init() {
	this.pool = redis.NewPool(func() (redis.Conn, error) {
		return redis.Connect(config.GetConfig("redis_host"))
		//return redis.Dial("tcp", config.GetConfig("redis_host"))
	}, 20)
	this.pool.IdleTimeout = 45 * time.Second
	this.pool.MaxActive = 30
	this.pool.MaxIdle = 30
}

func (this *Ssdb) Do(args ...interface{}) ([]string, error) { //(command string, args ...interface{}) interface{} {
	conn := this.pool.Get()
	defer conn.Close()
	reply, err := conn.Do(args...) //command, args...)
	//logger.Error("Ssdb command reply %v error :%v\n", reply, err)
	return reply, err
	/*if err != nil {
		logger.Error("Ssdb command error :%v\n", err)
		return nil
	}
	return reply*/
}


