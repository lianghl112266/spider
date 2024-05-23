package persist

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"spider/src/spider/module"
	"spider/src/spider_distribute/config"
)

// If you do not use distributed crawlers, use this saver
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.USERNAME, config.PASSWD, config.DATABASE_ADDR, config.DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&module.Weather{})

	go func() {
		for {
			item := <-out
			fmt.Println(item)
			if it, ok := item.(module.Weather); ok {
				fmt.Println(it)
				go func() { db.Create(&it) }()
			}
		}
	}()

	return out
}
