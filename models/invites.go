package models

type Invite struct {
	ID     int    `json:"id"`
	School School `json:"school"`
	Driver Driver `json:"driver"`
	Status string `json:"status"`
}
