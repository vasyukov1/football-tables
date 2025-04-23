package model

type Team struct {
    ID   int    `gorm:"primaryKey"`
    Name string
}
