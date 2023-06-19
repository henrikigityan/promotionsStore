package model

import (
	"time"
)

type Promotion struct {
	Id             string  `bson:"_id"`
	Price           float64  `bson:"price,omitempty"`
	ExpirationDate   time.Time `bson:"expiration_date,omitempty"`
}
