package neo4jdb

import (
	"fmt"
	"github.com/meowalien/rabbitgather-article/conf"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var GlobalNeo4jDriver  neo4j.Driver

func CreateNeo4jConnection (dbconf conf.Neo4JConfiguration)(neo4j.Driver,error){
	dbUri := "neo4j://%s:%s"
	dbUri = fmt.Sprintf(dbUri,dbconf.Host,dbconf.Port)
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth(dbconf.User, dbconf.Password, ""))
	if err != nil {
		return nil, err
	}
	return driver, err
}