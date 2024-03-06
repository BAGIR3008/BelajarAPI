package activity

type ActivityRequest struct {
	ID uint   `gorm:"primarykey"`
	Do string `json:"do" form:"do"`
}
