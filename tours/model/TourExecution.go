package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const _pointProximity = 100

type TourExecutionStatus int

const (
	Active TourExecutionStatus = iota
	Completed
	Abandoned
)

type TourExecution struct {
	Id        uuid.UUID `json:"id"`
	TourId    uuid.UUID `json:"toudId"`
	TouristId uuid.UUID `json:"touristId"`
	Status    TourExecutionStatus
	Position  Position `gorm:"type:jsonb"`
	Tasks     []PointTask
}

func (execution *TourExecution) BeforeCreate(scope *gorm.DB) error {
	execution.Id = uuid.New()
	return nil
}
