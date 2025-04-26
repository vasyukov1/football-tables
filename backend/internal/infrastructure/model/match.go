package model

type Match struct {
	ID          int `gorm:"primaryKey"`
	Team1ID     int
	Team2ID     int
	Score1      int    `gorm:"check:score1 >= 0"`
	Score2      int    `gorm:"check:score2 >= 0"`
	Stage       string `gorm:"type:varchar(50)"`
	GroupID     int
	PlayoffID   int
	IsCompleted bool `gorm:"not null;default:false"`
}
