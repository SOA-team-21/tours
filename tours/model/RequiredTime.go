package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransportType int

const (
	Walk TransportType = iota
	Drive
	Bicycle
)

type RequiredTime struct {
	Id        uuid.UUID     `json:"id" gorm:"primaryKey"`
	TourId    uuid.UUID     `json:"tourId"`
	Transport TransportType `json:"transport"`
	Minutes   int           `json:"minutes"`
}

func (requiredTime *RequiredTime) BeforeCreate(scope *gorm.DB) error {
	if err := requiredTime.Validate(); err != nil {
		return err
	}
	requiredTime.Id = uuid.New()
	return nil
}

func (requiredTime *RequiredTime) Validate() error {
	if requiredTime.Minutes > 60 || requiredTime.Minutes < 0 {
		return errors.New("invalid minutes")
	}
	return nil
}
