package mariadb

import (
	"fmt"
	"github.com/meowalien/rabbitgather-article/sec/conf"
	"github.com/meowalien/rabbitgather-lib/db_connect"
	"gorm.io/gorm"
)

var GlobalConn *gorm.DB
func InitMariadbDB() {
	fmt.Println("InitMariadbDB ...")

	var err error
	GlobalConn, err = db_connect.CreateGormDBConnection(conf.GlobalConfig.DB.MarinaDB)
	if err != nil {
		panic(fmt.Sprint("error when open MarinaDB connection with: ", conf.GlobalConfig.DB.MarinaDB, "error msg: ", err.Error()))
	}
}