package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/meowalien/rabbitgather-article/conf"
)
var GlobalRedisConn *redis.Client
func CreateRedisConnection(dbconf conf.RedisConfiguration) (*redis.Client , error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dbconf.Host + ":" + dbconf.Port,
		Password: dbconf.Password,
		DB:       dbconf.ID,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, err
}

