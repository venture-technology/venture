package entity

import "time"

type Partner struct {
	Record    int       `json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at"`
}
