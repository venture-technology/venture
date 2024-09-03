package entity

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Email struct {
	ID        uuid.UUID `json:"id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	Recipient string    `json:"recipient" validate:"required" example:"example@gmail.com"`
	Subject   string    `json:"subject" validate:"required" example:"subject - create account"`
	Body      string    `json:"body" validate:"required" example:"hello sr...e"`
}

func (e *Email) EmailStructToJson() (string, error) {
	json, err := json.Marshal(e)

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (e *Email) Unserialize(msg *kafka.Message) (*Email, error) {
	var email *Email

	err := json.Unmarshal(msg.Value, &email)
	if err != nil {
		log.Fatalf("Erro ao desserializar mensagem JSON: %v", err)
		return nil, err
	}

	return email, nil
}
