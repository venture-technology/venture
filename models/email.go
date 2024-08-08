package models

import "encoding/json"

type Email struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func (e *Email) EmailStructToJson() (string, error) {
	json, err := json.Marshal(e)

	if err != nil {
		return "", err
	}

	return string(json), nil
}
