package redis

import "github.com/go-redis/redis/v8"

func RedisConnection(redisAddress string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
