package model

import "github.com/google/uuid"

type TourPurchaseToken struct {
	TourId    uuid.UUID
	TouristId int64
}
