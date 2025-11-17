package orm

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	Title       string
	Description string     `gorm:"type:text"`
	CreatedByID uint
	IsAnonymous bool
	IsOpen      bool `gorm:"default:true"`

	Options []VoteOption `gorm:"constraint:OnDelete:CASCADE;"`
	Results []VoteResult `gorm:"constraint:OnDelete:CASCADE;"`
}

type VoteOption struct {
	gorm.Model
	VoteID uint   `gorm:"index"`
	Text   string
}

type VoteResult struct {
	gorm.Model
	VoteID   uint      `gorm:"index"`
	OptionID uint      `gorm:"index"`
	UserID   *uint     `gorm:"index"` // NULL when anonymous
}
