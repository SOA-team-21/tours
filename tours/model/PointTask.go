package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PointTask struct {
	Id              int64 `json:"id" gorm:"primaryKey"`
	TourExecutionId int64 `json:"tourExeId"`
	KeyPointId      int64 `json:"keyPointId"`
	Done            bool
	DoneOn          time.Time
}

func (task *PointTask) BeforeCreate(scope *gorm.DB) error {
	task.Id = int64(uuid.New().ID()) + time.Now().UnixNano()/int64(time.Microsecond)
	return nil
}
