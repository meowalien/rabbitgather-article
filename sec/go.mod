module github.com/meowalien/rabbitgather-article

go 1.16

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-redis/redis/v8 v8.11.3
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/meowalien/rabbitgather-lib v0.0.0-00010101000000-000000000000
	github.com/mitchellh/mapstructure v1.4.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/neo4j/neo4j-go-driver v1.8.3
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.9.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/sys v0.0.0-20210923061019-b8560ed6a9b7 // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)

replace github.com/meowalien/rabbitgather-lib => ../../lib
