package orm

import (
	"kitties/orm"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string

	Name       string
	Surname    string
	StudyGroup string
	Role       Role `gorm:"type:text"`

	Active bool `gorm:"default:true"`
}

// put the plaintext password inside the struct, it'll encrypt it and modify the struct
func (u *User) NewUser(db *gorm.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)

	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}
