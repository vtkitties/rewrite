package orm

import (
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
