package conf

import (
	"fmt"
	"github.com/kr/pretty"
	"github.com/meowalien/rabbitgather-lib/db_connect"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)


var DEBUG_MOD bool

func ReadConfig(configFile string) (ArticleConfig ,error){
	viper.SetConfigFile(configFile)
	err:= viper.ReadInConfig()
	if err != nil {
		return ArticleConfig{},err
	}
	settingMap := viper.AllSettings()
	articleConfig := ArticleConfig{}
	err = mapstructure.Decode(settingMap,&articleConfig)
	return articleConfig,err
}

func InitConfig(configFile string) {
	fmt.Println("InitConfig ...")
	//p,_:=filepath.Abs(config_file)
	var err error
	GlobalConfig, err = ReadConfig(configFile)
	if err != nil {
		panic(fmt.Sprint("error when open read config file: ", configFile))
	}
	if DEBUG_MOD {
		pretty.Println("GlobalConfig: ", GlobalConfig)
	}
}

var GlobalConfig ArticleConfig

// ArticleConfig is the root config struct
type ArticleConfig struct {
	DB      DatabaseConfig
	Servers ServerConfig
}

type ServerConfig struct {
	RestfulServer RestfulServerConfiguration
}

type RestfulServerConfiguration struct {
	Port string
}

type DatabaseConfig struct {
	MarinaDB db_connect.MysqlConnectConfiguration
	Neo4J db_connect.Neo4JConfiguration
	Redis db_connect.RedisConfiguration
}

