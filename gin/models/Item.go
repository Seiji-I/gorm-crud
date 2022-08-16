package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string	`gorm:"unique"`
}

/*
equal

type Item struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Name string	`gorm:"unique"`
}
*/