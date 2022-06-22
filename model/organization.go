package model

import "time"

type Organization struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	AvatarURL string     `json:"avatar_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
