package datasource

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"jhgocms/config"
	"time"
)

func NewRedis() {

}

/**
	Redis连接池
 */
func RedisPool() *redis.Pool  {
	var ReisPool *redis.Pool
	timeout := time.Duration(config.SysConfig.Redis.Timeout)*time.Second
	ReisPool = &redis.Pool{
		MaxIdle:config.SysConfig.Redis.MaxIdle,
		MaxActive:config.SysConfig.Redis.MaxActive,
		IdleTimeout:timeout,
		Wait:true,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", fmt.Sprintf("%s:%s",config.SysConfig.Redis.Host,config.SysConfig.Redis.Port),
				redis.DialPassword(config.SysConfig.Redis.Password),
				redis.DialDatabase(config.SysConfig.Redis.Database),
				redis.DialConnectTimeout(timeout),
				redis.DialReadTimeout(timeout),
				redis.DialWriteTimeout(timeout),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	return ReisPool
}