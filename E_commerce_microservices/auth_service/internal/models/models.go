package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	FirstName    string `gorm:"size:100"`
	LastName     string `gorm:"size:100"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
