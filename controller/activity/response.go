package activity

type ActivityResponse struct {
	ID uint   `gorm:"primarykey"`
	Do string `json:"do" form:"do"`
}
