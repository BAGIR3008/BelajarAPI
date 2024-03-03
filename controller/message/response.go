package message

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

func Response(status int, message string, data ...map[string]any) (int, map[string]any) {
	json := map[string]any{
		"code":    status,
		"message": message,
	}

	for _, part := range data {
		for key, value := range part {
			json[key] = value
		}
	}
	return status, json
}
