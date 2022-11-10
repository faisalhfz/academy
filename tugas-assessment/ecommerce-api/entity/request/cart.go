package request

import "ecommerce-api/entity"

type CreateCartRequest struct {
	Products   []entity.Product `json:"products" gorm:"many2many:cart_products"`
	TotalPrice int              `json:"totalprice"`
	IsCheckout bool             `json:"ischeckout"`
}

type TransactionPinRequest struct {
	PIN string `json:"pin"`
}
