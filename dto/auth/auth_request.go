package authDto

type AuthRequest struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" form:"name" gorm:"not null" validate:"required"`
	Email       string `json:"email" form:"email" gorm:"uniqueIndex;not null" validate:"required"`
	Password    string `json:"password" form:"password" gorm:"not null" validate:"required"`
	Address     string `json:"address" form:"address" gorm:"not null" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" gorm:"not null" validate:"required"`
	ImageUrl    string `json:"image_url" form:"image_url"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
