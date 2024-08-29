package entity

import "time"

type Admin struct {
	Name       string
	Position   string
	Created_at time.Time
}
