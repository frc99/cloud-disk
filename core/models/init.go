package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"xorm.io/xorm"
)

// var Engine = Init()
var Rdb = InitRedis()

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine error:%v", err)
		return nil
	}
	return engine
}
func InitRedis() redis.Conn {
	client, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	return client
}
