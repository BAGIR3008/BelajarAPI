package activity

import "time"

type ActivityResponse struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Do        string `json:"do" form:"do"`
}

type UserResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email" gorm:"unique"`
}
