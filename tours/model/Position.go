package model

import (
	"errors"
	"math"
	"time"

	"gorm.io/gorm"
)

type Position struct {
	Latitude     float64
	Longitude    float64
	LastActivity time.Time
}

func (position *Position) BeforeCreate(scope *gorm.DB) error {
	if err := position.Validate(); err != nil {
		return err
	}
	return nil
}

func (position *Position) Validate() error {
	if math.Abs(position.Latitude) > 90 {
		return errors.New("invalid latitude")
	}
	if math.Abs(position.Longitude) > 180 {
		return errors.New("invalid longitude")
	}
	return nil
}
