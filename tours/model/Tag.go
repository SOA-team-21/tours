package model

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	Tag string
}

func (tag *Tag) BeforeCreate(scope *gorm.DB) error {

	return nil
}

func (tag *Tag) Validate() error {
	if tag.Tag == "" {
		return errors.New("invalid tag name")
	}
	return nil
}
