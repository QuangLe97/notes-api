package models

import "time"

type LogMsg struct {
	ID        int64  `gorm:"primary_key;auto_increment" json:"id"`
	Message   string `gorm:"type:text" json:"message"`
	CreatedAt time.Time
}
