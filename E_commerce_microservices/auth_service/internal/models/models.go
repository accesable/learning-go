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
	Roles        []Role `gorm:"many2many:user_roles"`
}

type Role struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	RoleName string `gorm:"unique;not null"`
}

type UserRole struct {
	UserID uint
	RoleID uint
}
