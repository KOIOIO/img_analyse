package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
