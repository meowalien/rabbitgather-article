package redisdb

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/meowalien/rabbitgather-article/sec/conf"
	"github.com/meowalien/rabbitgather-lib/db_connect"
)
var GlobalRedisConn *redis.Client

func InitRedis() {
	fmt.Println("InitRedis ...")

	var err error
	GlobalRedisConn, err = db_connect.CreateRedisConnection(conf.GlobalConfig.DB.Redis)
	if err != nil {
		panic(fmt.Sprint("error when open Redis connection with: ", conf.GlobalConfig.DB.Redis, "error msg: ", err.Error()))
	}

}