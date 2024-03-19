package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TourStatus int

const (
	Draft TourStatus = iota
	Published
	Archived
)

type Tour struct {
	Id            int64 `json:"id" gorm:"primaryKey"`
	Name          string
	Description   string
	Difficult     int64
	Price         float64
	Status        TourStatus
	AuthorId      int64 `json:"authorId"`
	Length        float64
	PublishTime   time.Time
	ArchiveTime   time.Time
	MyOwn         bool
	KeyPoints     []KeyPoint
	RequiredTimes []RequiredTime
	Tags          pq.StringArray `gorm:"type:text[];"`
}

func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	if err := tour.Validate(); err != nil {
		return err
	}
	tour.Id = int64(uuid.New().ID()) + time.Now().UnixNano()/int64(time.Microsecond)
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
