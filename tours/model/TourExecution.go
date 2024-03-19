package model

import (
	"time"

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
	Id        int64 `json:"id"`
	TourId    int64 `json:"toudId"`
	TouristId int64 `json:"touristId"`
	Status    TourExecutionStatus
	Position  Position `gorm:"type:jsonb"`
	Tasks     []PointTask
}

func (execution *TourExecution) BeforeCreate(scope *gorm.DB) error {
	execution.Id = int64(uuid.New().ID()) + time.Now().UnixNano()/int64(time.Microsecond)
	return nil
}
