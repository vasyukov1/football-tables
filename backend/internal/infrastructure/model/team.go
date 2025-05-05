package model

import "time"

type Team struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
