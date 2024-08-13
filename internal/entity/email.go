package entity

import "encoding/json"

type Email struct {
	Recipient string `json:"recipient" validate:"required" example:"example@gmail.com"`
	Subject   string `json:"subject" validate:"required" example:"subject - create account"`
	Body      string `json:"body" validate:"required" example:"hello sr...e"`
}

func (e *Email) EmailStructToJson() (string, error) {
	json, err := json.Marshal(e)

	if err != nil {
		return "", err
	}

	return string(json), nil
}
