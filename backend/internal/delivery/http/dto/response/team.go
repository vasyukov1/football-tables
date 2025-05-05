package response

import "time"

type Team struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Played int `json:"played"`
	Wins int `json:"wins"`
	Draws int `json:"draws"`
	Loses int `json:"loses"`
	GoalsFor int `json:"goals_for"`
	GoalsAgainst int `json:"goals_against"`
	Points int `json:"points"`
	CreatedAt time.Time `json:"created_at"`
}
