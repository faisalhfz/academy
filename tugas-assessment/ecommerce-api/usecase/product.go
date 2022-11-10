package usecase

import (
	"ecommerce-api/entity"
	"ecommerce-api/entity/request"
	"ecommerce-api/entity/response"
	"ecommerce-api/repository"

	"github.com/jinzhu/copier"
)

type IProductUsecase interface {
	CreateProduct(requestBody request.CreateProductRequest) error
	GetProducts(filter request.ProductFilterRequest) ([]response.GetProductResponse, error)
	GetProductById(id int) (*entity.Product, error)
	EditProductById(id int, requestBody request.CreateProductRequest) error
	DeleteProductById(id int) error
}

type ProductUsecase struct {
	pRepository repository.IProductRepository
}

func NewProductUsecase(pRepository repository.IProductRepository) *ProductUsecase {
	return &ProductUsecase{pRepository: pRepository}
}

func (pUsecase ProductUsecase) CreateProduct(requestBody request.CreateProductRequest) error {
	product := entity.Product{}
	copier.Copy(&product, &requestBody)

	if err := pUsecase.pRepository.CreateProduct(product); err != nil {
		return err
	}
	return nil
}

func (pUsecase ProductUsecase) GetProducts(filter request.ProductFilterRequest) ([]response.GetProductResponse, error) {
	products, err := pUsecase.pRepository.GetAllProducts(filter)
	if err != nil {
		return nil, nil
	}
	productResponse := []response.GetProductResponse{}
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (pUsecase ProductUsecase) GetProductById(id int) (*entity.Product, error) {
	product, err := pUsecase.pRepository.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pUsecase ProductUsecase) EditProductById(id int, requestBody request.CreateProductRequest) error {
	product, err := pUsecase.pRepository.GetProduct(id)
	if err != nil {
		return err
	}
	productNew := entity.Product{}
	copier.Copy(&productNew, &requestBody)
	if err := pUsecase.pRepository.UpdateProduct(product, productNew); err != nil {
		return err
	}
	return nil
}

func (pUsecase ProductUsecase) DeleteProductById(id int) error {
	product, err := pUsecase.pRepository.GetProduct(id)
	if err != nil {
		return err
	}
	if err := pUsecase.pRepository.DeleteProduct(product); err != nil {
		return err
	}
	return nil
}
