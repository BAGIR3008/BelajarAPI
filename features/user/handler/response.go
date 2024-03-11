package handler

import "time"

type LoginResponse struct {
	Name  string
	Email string
	Token string
}

type ProfileResponse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Hp        string
}
