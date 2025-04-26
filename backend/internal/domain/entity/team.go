package entity

import "time"

type Team struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Barcelona"`
	CreatedAt time.Time `json:"created_at"`
}
