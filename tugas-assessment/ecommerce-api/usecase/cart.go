package usecase

import (
	"ecommerce-api/entity"
	"ecommerce-api/entity/request"
	"ecommerce-api/entity/response"
	"ecommerce-api/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	CreateCart(requestBody request.CreateCartRequest) error
	GetCart() (*entity.Cart, error)
	AddProductToCart(newProduct entity.Product) error
	CheckoutCart() error
	GetCompletedCarts() ([]response.GetCartResponse, error)
	ClearCart() error
	// RemoveProductFromCart(productId int) error
	// IsProductInCart(productId int, cart *entity.Cart) bool
}

type CartUsecase struct {
	cRepository repository.ICartRepository
}

func NewCartUsecase(cRepository repository.ICartRepository) *CartUsecase {
	return &CartUsecase{cRepository: cRepository}
}

func (cUsecase CartUsecase) CreateCart(requestBody request.CreateCartRequest) error {
	cart := entity.Cart{}
	copier.Copy(&cart, &requestBody)
	if err := cUsecase.cRepository.CreateCart(cart); err != nil {
		return err
	}
	return nil
}

func (cUsecase CartUsecase) GetCart() (*entity.Cart, error) {
	cart, err := cUsecase.cRepository.GetCart()
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (cUsecase CartUsecase) AddProductToCart(newProduct entity.Product) error {
	if err := cUsecase.cRepository.AddProduct(newProduct); err != nil {
		return err
	}
	return nil
}

func (cUsecase CartUsecase) CheckoutCart() error {
	if err := cUsecase.cRepository.CheckoutCart(); err != nil {
		return err
	}
	return nil
}

func (cUsecase CartUsecase) GetCompletedCarts() ([]response.GetCartResponse, error) {
	carts, err := cUsecase.cRepository.GetCompletedCarts()
	if err != nil {
		return nil, err
	}
	cartsResponse := []response.GetCartResponse{}
	copier.Copy(&cartsResponse, &carts)
	return cartsResponse, nil
}

func (cUsecase CartUsecase) ClearCart() error {
	cart, err := cUsecase.cRepository.GetCart()
	if err != nil {
		return err
	}
	if err := cUsecase.cRepository.DeleteCart(cart); err != nil {
		return err
	}
	return nil
}

// func (cUsecase CartUsecase) RemoveProductFromCart(productId int) error {
// 	if err := cUsecase.cRepository.RemoveProduct(productId); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (cUsecase CartUsecase) IsProductInCart(productId int, cart *entity.Cart) bool {
// 	productList := cart.Products
// 	for _, product := range productList {
// 		if product.ID == productId {
// 			return true
// 		}
// 	}
// 	return false
// }
