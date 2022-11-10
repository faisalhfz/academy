package entity

type Cart struct {
	ID         int       `json:"id" gorm:"column:id;type:bigint;primaryKey;autoincrement"`
	Products   []Product `json:"products" gorm:"many2many:cart_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TotalPrice int       `json:"totalprice"`
	IsCheckout bool      `json:"ischeckout"`
}
