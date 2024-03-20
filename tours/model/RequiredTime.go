package model

import (
	"errors"

	"gorm.io/gorm"
)

type TransportType int

const (
	Walk TransportType = iota
	Drive
	Bicycle
)

type RequiredTime struct {
	Id        int64         `json:"id" gorm:"primaryKey"`
	TourId    int64         `json:"tourId"`
	Transport TransportType `json:"transport"`
	Minutes   int           `json:"minutes"`
}

func (requiredTime *RequiredTime) BeforeCreate(scope *gorm.DB) error {
	if err := requiredTime.Validate(); err != nil {
		return err
	}
	return nil
}

func (requiredTime *RequiredTime) Validate() error {
	if requiredTime.Minutes < 0 {
		return errors.New("invalid minutes")
	}
	return nil
}
