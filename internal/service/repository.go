package service

import "github.com/LeUrok/DS-lab1/internal/model"

type PersonRepository interface {
	Create(req model.PersonRequest) (int32, error)
	GetByID(id int32) (*model.Person, error)
	GetAll() ([]model.Person, error)
	Update(id int32, req model.PersonRequest) (*model.Person, error)
	Delete(id int32) error
}