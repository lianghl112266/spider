package persist

import (
	"fmt"
	"gorm.io/gorm"
	"spider/module"
)

type ItemSaverService struct {
	DB *gorm.DB
}

// Save called by the client
func (this *ItemSaverService) Save(item module.Weather, _ *string) error {
	this.DB.Create(&item)
	fmt.Println(item)
	return nil
}
