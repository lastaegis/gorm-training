package Models

import (
	"database/sql"
	"time"
)

// User models digunakan untuk melakukan eksekusi query ke database
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt sql.NullTime
	DeletedBy sql.NullString
}
