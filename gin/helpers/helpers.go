package helpers

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
	models "gin/models"
)

func GetAllItems() []models.Item{
	dsn := "host=postgres user=root password=root dbname=sample port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	var items []models.Item
	db.Find(&items)
	return items
}

func CreateOneItem(input_item string) bool{
	dsn := "host=postgres user=root password=root dbname=sample port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err!= nil {
		log.Fatal(err)
		return false
	}
	tx := db.Begin()
	item := models.Item{Name: input_item}
	result := tx.Create(&item)
	tx.Save(&item)
	if result.Error != nil {
		log.Fatal(result.Error)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}