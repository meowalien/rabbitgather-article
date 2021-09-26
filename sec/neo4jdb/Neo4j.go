package neo4jdb

import (
	"fmt"
	"github.com/meowalien/rabbitgather-article/sec/conf"
	"github.com/meowalien/rabbitgather-lib/db_connect"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var GlobalNeo4jDriver  neo4j.Driver

func InitNeo4j() {
	fmt.Println("InitNeo4j ...")

	var err error
	GlobalNeo4jDriver, err = db_connect.CreateNeo4jConnection(conf.GlobalConfig.DB.Neo4J)
	if err != nil {
		panic(fmt.Sprint("error when open Neo4j connection with: ", conf.GlobalConfig.DB.Neo4J, "error msg: ", err.Error()))
	}

}