package models

import (
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Duration `gorm:"column:create_at" json:"-"`
	UpdatedAt time.Duration `gorm:"column:update_at" json:"-"`
}

type Token struct {
	Token string `json:"token"`
}
