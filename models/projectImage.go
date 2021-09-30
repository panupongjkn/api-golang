package models

import "github.com/jinzhu/gorm"

type ProjectImage struct {
	gorm.Model
	ImageURL  string `gorm:"not null"`
	ProjectID uint
}
