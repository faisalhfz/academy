package repository

import (
	"ecommerce-api/entity"
	"ecommerce-api/entity/request"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product entity.Product) error
	GetAllProducts(filter request.ProductFilterRequest) ([]entity.Product, error)
	GetProduct(id int) (*entity.Product, error)
	UpdateProduct(product *entity.Product, productNew entity.Product) error
	DeleteProduct(product *entity.Product) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pRepository ProductRepository) CreateProduct(product entity.Product) error {
	if err := pRepository.db.Debug().Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (pRepository ProductRepository) GetAllProducts(filter request.ProductFilterRequest) ([]entity.Product, error) {
	var products []entity.Product

	category := filter.Category
	min := filter.MinPrice
	max := filter.MaxPrice

	if category != "" && max != 0 {
		if err := pRepository.db.Where("category = ?", category).Where("price >= (?)", min).Where("price <= (?)", max).Find(&products).Error; err != nil {
			return nil, err
		}
		return products, nil
	}
	if category != "" && max == 0 {
		if err := pRepository.db.Where("category = ?", category).Where("price >= (?)", min).Find(&products).Error; err != nil {
			return nil, err
		}
		return products, nil
	}
	if category == "" && max != 0 {
		if err := pRepository.db.Where("price >= (?)", min).Where("price <= (?)", max).Find(&products).Error; err != nil {
			return nil, err
		}
		return products, nil
	}
	if err := pRepository.db.Where("price >= (?)", min).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pRepository ProductRepository) GetProduct(id int) (*entity.Product, error) {
	var product *entity.Product
	if err := pRepository.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (pRepository ProductRepository) UpdateProduct(product *entity.Product, productNew entity.Product) error {
	if err := pRepository.db.Debug().Model(&product).Updates(productNew).Error; err != nil {
		return err
	}
	return nil
}

func (pRepository ProductRepository) DeleteProduct(product *entity.Product) error {
	if err := pRepository.db.Debug().Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
