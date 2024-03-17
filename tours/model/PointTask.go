package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PointTask struct {
	Id              uuid.UUID `json:"id" gorm:"primaryKey"`
	TourExecutionId uuid.UUID `json:"tourExeId"`
	KeyPointId      uuid.UUID `json:"keyPointId"`
	Done            bool
	DoneOn          time.Time
}

func (task *PointTask) BeforeCreate(scope *gorm.DB) error {
	task.Id = uuid.New()
	return nil
}
