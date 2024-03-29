package Migrations

import (
	"database/sql"
	"time"
)

// User migration digunakan untuk melakukan pembentukan table pada database
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"index;size:255"`
	Email     string `gorm:"uniqueIndex;size:255;comment:Email yang digunakan bersifat unique"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt sql.NullTime
	DeletedBy sql.NullString
}
