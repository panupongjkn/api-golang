package models

import "github.com/jinzhu/gorm"

type ProjectTag struct {
	gorm.Model
	Name      string `gorm:"not null"`
	ProjectID uint
}
