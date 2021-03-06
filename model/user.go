package model

import "time"

type User struct {
	ID        string     `json:"id"`
	Login     string     `json:"login"`
	Password  []byte     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
