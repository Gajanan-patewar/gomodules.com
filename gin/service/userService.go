package service

import "example/curd/model"

type UserService interface {
	CreateUser(*model.User) error
	GetUser(*string) (*model.User, error)
	GetAll() ([]*model.User, error)
	UpdateUser(*model.User) error
	DeleteUser(*string) error
}
