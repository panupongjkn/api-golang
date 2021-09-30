package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name     string `gorm:"unique; not null"`
	Cover    string `gorm:"not null"`
	Like     int    `gorm:"not null"`
	Download int    `gorm:"not null"`
	Program  string `gorm:"not null"`
	OwnerID  uint

	ProjectImages []ProjectImage `gorm:"foreignKey:ProjectID"`
	ProjectTags   []ProjectTag   `gorm:"foreignKey:ProjectID"`
}
