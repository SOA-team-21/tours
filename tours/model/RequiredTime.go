package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TransportType int

const (
	Walk TransportType = iota
	Drive
	Bicycle
)

type RequiredTime struct {
	Transport TransportType
	Minutes   int
}

func (requiredTime *RequiredTime) BeforeCreate(scope *gorm.DB) error {
	if err := requiredTime.Validate(); err != nil {
		return err
	}
	return nil
}

func (requiredTime *RequiredTime) Validate() error {
	if requiredTime.Minutes > 60 || requiredTime.Minutes < 0 {
		return errors.New("invalid minutes")
	}
	return nil
}

func (d *RequiredTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, d)
}

func (d RequiredTime) Value() (driver.Value, error) {
	return json.Marshal(d)
}
