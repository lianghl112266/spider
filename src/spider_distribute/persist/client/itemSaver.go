package client

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"spider/src/spider/module"
	"spider/src/spider_distribute/config"
	"spider/src/spider_distribute/rpc_support"
)

func ItemSaver(host string) chan interface{} {
	out := make(chan interface{})
	client, err := rpc_support.NewClient(host)
	if err != nil {
		fmt.Println(err)
		return out
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.USERNAME, config.PASSWD, config.DATABASE_ADDR, config.DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&module.Weather{})

	go func() {
		for {
			item := <-out
			if it, ok := item.(module.Weather); ok {
				go func() { _ = client.Call("ItemSaverService.Save", it, "") }()
			}
		}
	}()

	return out
}
