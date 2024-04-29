package database

import "time"

type User struct {
	Email      string    `json:"email,omitempty"`
	Name       string    `json:"name,omitempty"`
	HashedPass string    `json:"hashed_pass,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
