package model

import (
	"gorm.io/gorm"
)

type Preference struct {
	Id         int64 `json:"id" gorm:"primaryKey"`
	UserId     int64
	Transport  string
	Difficulty int
	Tags       string
}

func (task *Preference) BeforeCreate(scope *gorm.DB) error {
	return nil
}
