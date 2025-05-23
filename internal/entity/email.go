package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Email struct {
	ID        uuid.UUID `json:"id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	Recipient string    `json:"recipient" validate:"required" example:"example@gmail.com"`
	Subject   string    `json:"subject" validate:"required" example:"subject - create account"`
	Body      string    `json:"body" validate:"required" example:"hello sr...e"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (e *Email) Serialize() (string, error) {
	json, err := json.Marshal(e)

	if err != nil {
		return "", err
	}

	return string(json), nil
}
