package conf

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

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

var GlobalConfig ArticleConfig

// ArticleConfig is the root config struct
type ArticleConfig struct {
	DB DatabaseConfig
	Handler HandlerConfig
}

type HandlerConfig struct {
	RestfulHandler RestfulHandlerConfiguration
}

type RestfulHandlerConfiguration struct {
	Port string
}

type DatabaseConfig struct {
	MarinaDB DatabaseConnectConfiguration
	Neo4J    Neo4JConfiguration
	Redis RedisConfiguration
}

type DatabaseConnectConfiguration struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

type Neo4JConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
}

type RedisConfiguration struct {
	Host string
	Port string
	Password string
	ID int
}