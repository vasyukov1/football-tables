package model

type Stage struct {
	ID        int `gorm:"primaryKey"`
	PlayoffID *int
	Name      string
	Matches   []*Match `gorm:"foreignKey:Stage"`
}
