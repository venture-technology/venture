package entity

import "github.com/google/uuid"

type Invite struct {
	ID     uuid.UUID `json:"id,omitempty"`
	School School    `json:"school,omitempty"`
	Driver Driver    `json:"driver,omitempty"`
	Status string    `json:"status,omitempty"`
}
