package engine

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisHelper struct {
	*redis.Client
}

var RedisClient *redis.Client
var redisHelper *RedisHelper

var redisOnce sync.Once

func GetRedisHelper() *RedisHelper {
	return redisHelper
}

func init() {
	RedisClient = NewRedisHelper()
}

func NewRedisHelper() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "172.16.6.92:6379",
		Password:     "winex",
		DB:           1,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})

	return rdb
}
