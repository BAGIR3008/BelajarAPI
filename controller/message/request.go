package message

type LoginRequest struct {
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=4,max=20,alpha"`
	Email    string `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,alphanumunicode"`
}
