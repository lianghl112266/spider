package main

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"spider/config"
	"spider/module"
	"spider/persist"
	"spider/rpcBase"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err.Error())
	}
}

// Persistence server function entry
func main() {
	dsn := viper.GetString("db.dsn")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = db.AutoMigrate(&module.Weather{})
	_ = rpcBase.ServeRPC(":1235", &persist.ItemSaverService{DB: db})
}
