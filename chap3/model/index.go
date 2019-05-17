package model

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	ID        uint      `gorm:"id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
