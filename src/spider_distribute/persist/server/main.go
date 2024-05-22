package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"spider/src/spider_distribute/config"
	"spider/src/spider_distribute/persist"
	"spider/src/spider_distribute/rpc_support"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.USERNAME, config.PASSWD, config.DATABASE)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = rpc_support.ServeRPC(config.HOST, &persist.ItemSaverService{DB: db})
}
