package core

import (
	"strconv"
	"github.com/go-redis/redis"
)

type RedisConnection struct {
	Config FileContent
	Client *redis.Client
}

func NewRedisConnection() CacheConnection{
	return &RedisConnection{}
}

func (r *RedisConnection) initCacheConnection(content FileContent) CacheConnection {
	r.Config = content
	address := content["host"].(string) + ":" + strconv.Itoa(content["port"].(int))
	pwd, ok := content["password"].(string)

	if  !ok {
		pwd = ""
	}

	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: pwd,
		MinIdleConns: content["min-idle-conns"].(int),
		PoolSize: content["pool-size"].(int),
	})

	_, err := client.Ping().Result()

	if nil != err {
		panic("cache connect failed")
	}

	r.Client = client

	return r
}