package repository

import (
	"user-management/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll(qname string) ([]entity.User, error)
	GetById(id int) (*entity.User, error)
	Update(user *entity.User, userNew entity.User) error
	Delete(user *entity.User, id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Debug().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAll(qname string) ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Where("name LIKE ?", qname+"%").Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

func (u UserRepository) GetById(id int) (*entity.User, error) {
	var user *entity.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) Update(user *entity.User, userNew entity.User) error {
	if err := u.db.Debug().Model(&user).Updates(userNew).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Delete(user *entity.User, id int) error {
	if err := u.db.Debug().Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
