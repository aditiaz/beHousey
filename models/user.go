package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ListAs    string    `json:"listAs"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address" gorm:"type: text"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
