package service

import (
	"errors"
	"github.com/LeUrok/DS-lab1/internal/model"
)

var (
	ErrNotFound = errors.New("person not found")
	ErrInvalidInput = errors.New("invalid input")
)

type PersonService struct {
	rep PersonRepository
}

func NewPersonService (rep PersonRepository) *PersonService {
	return &PersonService{rep: rep}
}

func (s *PersonService) Create(req model.PersonRequest) (int32, error) {
	if req.Name == "" {
		return 0, ErrInvalidInput
	}

	return s.rep.Create(req)
}

func (s *PersonService) GetByID(id int32) (*model.Person, error) {
	if id <= 0 {
		return nil, ErrInvalidInput
	}

	return s.rep.GetByID(id)
}

func (s *PersonService) GetAll() ([]model.Person, error) {
	return s.rep.GetAll()
}

func (s *PersonService) Update(id int32, req model.PersonRequest) (*model.Person, error) {
	if id <= 0 {
		return nil, ErrInvalidInput
	}
	if req.Address == "" && req.Age == 0 && req.Name == "" && req.Work == "" {
		return nil, ErrInvalidInput
	}
	return s.rep.Update(id, req)
}

func (s *PersonService) Delete(id int32) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	return s.rep.Delete(id)
}