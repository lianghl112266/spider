package persist

import (
	"fmt"
	"gorm.io/gorm"
	"spider/src/spider/module"
)

type ItemSaverService struct {
	DB *gorm.DB
}

// The remote function called by the client is this
func (this *ItemSaverService) Save(item module.Weather, _ *string) error {
	this.DB.Create(&item)
	fmt.Println(item)
	return nil
}
