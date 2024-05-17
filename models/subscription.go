package models

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
}
