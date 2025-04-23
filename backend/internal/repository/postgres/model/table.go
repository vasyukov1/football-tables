package model

type Table struct {
    ID       int   `gorm:"primaryKey"`
    GroupIDs []int `gorm:"-"`
}
