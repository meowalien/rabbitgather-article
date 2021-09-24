package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kr/pretty"
	"github.com/meowalien/rabbitgather-article/conf"
	"github.com/meowalien/rabbitgather-article/global"
	"github.com/meowalien/rabbitgather-article/mariadb"
	"github.com/meowalien/rabbitgather-article/neo4jdb"
	"github.com/meowalien/rabbitgather-article/redisdb"
	"github.com/meowalien/rabbitgather-article/restful"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var DEBUGMOD = false

const config_file = "../config/config.json"

func initFlags() {
	flag.BoolVar(&DEBUGMOD, "debug", false, "\"true\" to open debug mode")
	flag.Parse()

	fmt.Println("debug mod: ", DEBUGMOD)
}

func initConfig() {
	var err error
	conf.GlobalConfig, err = conf.ReadConfig(config_file)
	if err != nil {
		log.Fatal("error when open read config file: ", config_file)
	}
	if DEBUGMOD {
		pretty.Println("GlobalConfig: ", conf.GlobalConfig)
	}
}

func initDB() {
	var err error
	mariadb.GlobalSQLConn, err = mariadb.CreateDBConnection(conf.GlobalConfig.DB.MarinaDB)
	if err != nil {
		log.Fatal("error when open MarinaDB connection with: ", conf.GlobalConfig.DB.MarinaDB, "error msg: ", err.Error())
	}
	addToFinalize(func() {
		e := mariadb.CloseConn(mariadb.GlobalSQLConn)
		if e != nil {
			fmt.Println("error when cosing GlobalSQLConn: ", e.Error())
		}
	})
}

func initNeo4j() {
	var err error
	neo4jdb.GlobalNeo4jDriver, err = neo4jdb.CreateNeo4jConnection(conf.GlobalConfig.DB.Neo4J)
	if err != nil {
		log.Fatal("error when open Neo4j connection with: ", conf.GlobalConfig.DB.Neo4J, "error msg: ", err.Error())
	}
	addToFinalize(func() {
		e := neo4jdb.GlobalNeo4jDriver.Close()
		if e != nil {
			fmt.Println("error when cosing GlobalNeo4jDriver: ", e.Error())
		}
	})
}

func initRedis() {
	var err error
	redisdb.GlobalRedisConn, err = redisdb.CreateRedisConnection(conf.GlobalConfig.DB.Redis)
	if err != nil {
		log.Fatal("error when open Redis connection with: ", conf.GlobalConfig.DB.Redis, "error msg: ", err.Error())
	}
	addToFinalize(func() {
		e := redisdb.GlobalRedisConn.Close()
		if e != nil {
			fmt.Println("error when cosing GlobalRedisConn: ", e.Error())
		}
	})
}

func initLogger() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		// time格式
		TimestampFormat: time.StampNano,
		PrettyPrint:     true,
	})
	logger.SetReportCaller(true)
	//輸出終端機
	logger.SetOutput(io.MultiWriter(os.Stdout))
	//設定log等級
	logger.SetLevel(logrus.DebugLevel)
	entry := logger.WithFields(logrus.Fields{
		"server": "the server name",
	})
	global.Logger = entry
}

func init() {
	initFlags()
	initConfig()
	initDB()
	initNeo4j()
	initRedis()
	initLogger()
}

func main() {
	defer finalize()
	global.Logger.Infof("theiohdkjljfdslkfjsdlfd")
	ctx := context.Background()

	restfulHandler := restful.RestfulHandler{
		Debug: DEBUGMOD,
	}
	go restfulHandler.Start(ctx)

	waitForShutdown(ctx, func() {
		var err error
		err = restfulHandler.Stop(ctx)
		if err != nil {
			fmt.Println("Error when close RestfulHandler: ", err.Error())
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
