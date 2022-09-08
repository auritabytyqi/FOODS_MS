package storage

import (
	"log"

	config "FOODS_MS/config"
	"FOODS_MS/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func AddFoodRecord(food model.Foods) error {
	DB.Select("Id", "Name", "Description").Create(&food)
	return nil
}

func DeleteFood(id string) error {
	DB.Delete(&model.Foods{}, id)
	return nil
}

func UpdateFood(id, name, description string) error {
	DB.Model(&model.Foods{}).Where("id=?", id).Updates(model.Foods{Name: name, Description: description})
	return nil
}
