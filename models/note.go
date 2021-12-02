package models

import "time"

type Note struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Content   string `gorm:"type:text" json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notes []Note
