package driver

import (
	"gopkg.in/redis.v5"
)

// CreateRedis 初始化Redis组件
func CreateRedis(addr, pwd string, db int) (*redis.Client, error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisClient, nil
}
