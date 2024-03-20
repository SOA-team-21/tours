package model

import (
	"errors"
	"math"

	"gorm.io/gorm"
)

type KeyPoint struct {
	Id          int64 `json:"id" gorm:"primaryKey"`
	TourId      int64 `json:"tourId"`
	Latitude    float64
	Longitude   float64
	Name        string
	Description string
	Picture     string
	Public      bool
}

func (point *KeyPoint) BeforeCreate(scope *gorm.DB) error {
	if err := point.Validate(); err != nil {
		return err
	}
	return nil
}

func (point *KeyPoint) Validate() error {
	if point.Name == "" {
		return errors.New("invalid name")
	}
	if point.Description == "" {
		return errors.New("invalid description")
	}
	if point.Picture == "" {
		return errors.New("invalid picture")
	}
	if math.Abs(point.Latitude) > 90 {
		return errors.New("invalid latitude")
	}
	if math.Abs(point.Longitude) > 180 {
		return errors.New("invalid longitude")
	}
	return nil
}
