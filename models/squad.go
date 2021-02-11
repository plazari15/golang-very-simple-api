package models

import "time"

type Squad struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"names"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"deletedcvAt"`
}
