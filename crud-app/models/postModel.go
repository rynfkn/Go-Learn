package models

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Post struct {
	Model
	Title string
	Body  string
}
