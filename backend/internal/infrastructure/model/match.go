package model

// type Match struct {
// 	ID          int `gorm:"primaryKey"`
// 	Team1ID     int
// 	Team2ID     int
// 	Score1      int    `gorm:"check:score1 >= 0"`
// 	Score2      int    `gorm:"check:score2 >= 0"`
// 	Stage       string `gorm:"type:varchar(50)"`
// 	// GroupID     int
// 	// PlayoffID   int
// 	IsCompleted bool `gorm:"not null;default:false"`
	
// 	GroupID      *int
// 	PlayoffID    *int
// 	NextMatchID  *int
// }

type Match struct {
    ID           int    `gorm:"primaryKey"`
    Team1ID      int    `gorm:"not null"`
    Team2ID      int    `gorm:"not null"`
    Score1       int    `gorm:"check:score1 >= 0;default:0"`
    Score2       int    `gorm:"check:score2 >= 0;default:0"`
    Stage        string `gorm:"type:varchar(50);not null"`
    GroupID      *int
    PlayoffID    *int
    NextMatchID  *int
    IsCompleted  bool   `gorm:"not null;default:false"`
}