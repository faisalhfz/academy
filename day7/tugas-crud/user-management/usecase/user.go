package usecase

import (
	"user-management/entity"
	"user-management/entity/response"
	"user-management/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(req response.CreateUserRequest) error
	GetList(qname string) ([]response.GetUserResponse, error)
	GetUser(id int) (response.GetUserResponse, error)
	UpdateUser(id int, newReq response.CreateUserRequest) error
	DeleteUser(id int) error
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) GetList(qname string) ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll(qname)
	if err != nil {
		return nil, nil
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}

func (u UserUsecase) GetUser(id int) (*response.GetUserResponse, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	userResponse := &response.GetUserResponse{}
	copier.Copy(&userResponse, &user)
	return userResponse, nil
}

func (u UserUsecase) UpdateUser(id int, newReq response.CreateUserRequest) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return err
	}
	userNew := entity.User{}
	copier.Copy(&userNew, &newReq)
	if err := u.userRepository.Update(user, userNew); err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) DeleteUser(id int) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return err
	}
	if err := u.userRepository.Delete(user, id); err != nil {
		return err
	}
	return nil
}
