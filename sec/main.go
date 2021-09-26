package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/meowalien/rabbitgather-article/sec/conf"
	"github.com/meowalien/rabbitgather-article/sec/logger"
	"github.com/meowalien/rabbitgather-article/sec/mariadb"
	"github.com/meowalien/rabbitgather-article/sec/neo4jdb"
	"github.com/meowalien/rabbitgather-article/sec/redisdb"
	"github.com/meowalien/rabbitgather-article/sec/server"
	"github.com/meowalien/rabbitgather-lib/db_connect"
	"os"
	"os/signal"
	"syscall"
)

const config_file = "../config/config.json"

func init() {
	InitFlags()
	conf.InitConfig(config_file)
	mariadb.InitMariadbDB()
	addToFinalize(func() {
		e := db_connect.CloseConn(mariadb.GlobalConn)
		if e != nil {
			fmt.Println("error when cosing GlobalConn: ", e.Error())
		}
	})
	neo4jdb.InitNeo4j()
	addToFinalize(func() {
		e := neo4jdb.GlobalNeo4jDriver.Close()
		if e != nil {
			fmt.Println("error when cosing GlobalNeo4jDriver: ", e.Error())
		}
	})
	redisdb.InitRedis()
	addToFinalize(func() {
		e := redisdb.GlobalRedisConn.Close()
		if e != nil {
			fmt.Println("error when cosing GlobalRedisConn: ", e.Error())
		}
	})
	logger.InitLogger()
}
func InitFlags() {
	flag.BoolVar(&conf.DEBUG_MOD, "debug", false, "\"true\" to open debug mode")

	flag.Parse()

	fmt.Println("debug mod: ", conf.DEBUG_MOD)
}


func main() {
	defer finalize()
	ctx := context.Background()

	restfulHandler := server.Server{
		Config: conf.GlobalConfig.Servers.RestfulServer,
		Debug:  conf.DEBUG_MOD,
	}
	go restfulHandler.Start(ctx)

	waitForShutdown(ctx, func() {
		var err error
		err = restfulHandler.Stop(ctx)
		if err != nil {
			fmt.Println("Error when close Server: ", err.Error())
			err = nil
		}
	})
}

func waitForShutdown(ctx context.Context, callbackFunc func()) {
	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fmt.Println("Shutdown with Context done")
	case <-quitSignal:
		fmt.Println("Shutdown with OS QuitSignal")
		callbackFunc()
	}
}

var finalizeFuncList = []func(){}

func addToFinalize(f func()) {
	finalizeFuncList = append(finalizeFuncList, f)
}
func finalize() {
	for _, f := range finalizeFuncList {
		f()
	}
}
