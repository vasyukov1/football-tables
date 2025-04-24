package model

type Playoff struct {
	ID     int      `gorm:"primaryKey"`
	Rounds []*Stage `gorm:"foreignKey:PlayoffID"`
}
