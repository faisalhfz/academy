package repository

import (
	"ecommerce-api/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(newCart entity.Cart) error
	GetCart() (*entity.Cart, error)
	AddProduct(newProduct entity.Product) error
	CheckoutCart() error
	GetCompletedCarts() ([]entity.Cart, error)
	DeleteCart(cart *entity.Cart) error
	// RemoveProduct(productId int) error
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (cRepository CartRepository) CreateCart(cart entity.Cart) error {
	if err := cRepository.db.Debug().Create(&cart).Error; err != nil {
		return err
	}
	products := cart.Products
	cRepository.db.Model(&cart).Association("Products").Append(&products)
	return nil
}

func (cRepository CartRepository) GetCart() (*entity.Cart, error) {
	var cart *entity.Cart
	if err := cRepository.db.Preload("Products").Where("is_checkout = ?", false).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (cRepository CartRepository) AddProduct(newProduct entity.Product) error {
	cart, err := cRepository.GetCart()
	if err != nil {
		return err
	}
	cRepository.db.Preload("Products").First(&cart)
	cart.Products = append(cart.Products, newProduct)
	cart.TotalPrice += newProduct.Price
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (cRepository CartRepository) CheckoutCart() error {
	cart, err := cRepository.GetCart()
	if err != nil {
		return err
	}
	cRepository.db.Preload("Products").First(&cart)
	cart.IsCheckout = true
	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (cRepository CartRepository) GetCompletedCarts() ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := cRepository.db.Preload("Products").Where("is_checkout = ?", true).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (cRepository CartRepository) DeleteCart(cart *entity.Cart) error {
	var products []entity.Product
	cRepository.db.Model(entity.Product{}).Find(&products)
	cRepository.db.Model(&cart).Association("Languages").Delete(products)
	if err := cRepository.db.Debug().Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}

// func (cRepository CartRepository) RemoveProduct(productId int) error {
// 	cart, err := cRepository.GetCart()
// 	if err != nil {
// 		return err
// 	}
// 	productList := cart.Products
// 	var index int
// 	for i, product := range productList {
// 		if product.ID == productId {
// 			index = i
// 			break
// 		}
// 	}
// 	cRepository.db.Preload("Products").First(&cart)
// 	cart.TotalPrice -= cart.Products[index].Price
// 	cart.Products = append(cart.Products[:index], cart.Products[index+1:]...)
// 	if err := cRepository.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
