package productDto

type CreateProductRequest struct {
	Name        string  `json:"name" form:"name" gorm:"not null"`
	Description string  `json:"description" form:"description" gorm:"not null"`
	Price       float64 `json:"price" form:"price" gorm:"not null"`
	ImageURL    string  `json:"image_url" form:"image_url" gorm:"not null"`
	PartnerID   uint    `json:"partner_id" form:"partner_id" gorm:"not null"`
}
