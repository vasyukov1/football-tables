package model

type Group struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Teams   []*Team  `gorm:"many2many:group_teams"`
	Matches []*Match `gorm:"foreignKey:GroupID"`
}
