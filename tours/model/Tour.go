package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
)

type Tour struct {
	Id            uuid.UUID `json:"id" gorm:"primaryKey"`
	Name          string
	Description   string
	Difficult     int64
	Price         float64
	Status        TourStatus
	AuthorId      uuid.UUID `json:"authorId"`
	Length        float64
	PublishTime   time.Time
	ArchiveTime   time.Time
	MyOwn         bool
	KeyPoints     []KeyPoint
	RequiredTimes []RequiredTime `gorm:"type:jsonb;"`
	Tags          []Tag          `gorm:"type:jsonb;"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	if err := tour.Validate(); err != nil {
		return err
	}
	tour.Id = uuid.New()
	return nil
}

func (tour *Tour) Validate() error {
	if tour.Name == "" {
		return errors.New("invalid name")
	}
	if tour.Description == "" {
		return errors.New("invalid description")
	}
	if tour.Difficult < 0 {
		return errors.New("invalid difficult, should be above 0")
	}
	if tour.Price < 0.0 {
		return errors.New("invalid price, should be above 0")
	}
	return nil
}
