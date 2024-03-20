package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

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
	Position  Position `gorm:"type:jsonb;"`
	Tasks     []PointTask
}

func (execution *TourExecution) BeforeCreate(scope *gorm.DB) error {
	return nil
}

func (a Position) Value() (driver.Value, error) {
	return json.Marshal(a)
}
func (a *Position) Scan(value interface{}) error {
	if value == nil {
		*a = Position{}
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}
