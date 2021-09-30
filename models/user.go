package models

import (
	"log"

	_ "github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique; not null"`
	Password     string `gorm:"not null"`
	FirstName    string `gorm:"not null" json:"first_name"`
	LastName     string `gorm:"not null" json:"last_name"`
	ProfileImage string

	Projects []Project `gorm:"foreignKey:OwnerID"`
}

func (u *User) CreateUser() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	orm.Create(u)
}

func (u *User) GetUserByEmail() error {
	orm.Where("email = ?", u.Email).First(u)
	return nil
}

func (u *User) GetUserByID() error {
	orm.Where("id = ?", u.ID).First(u)
	return nil
}
