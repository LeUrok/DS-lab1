package service

import (
	"errors"
	"testing"

	"github.com/LeUrok/DS-lab1/internal/model"
)

type mockRepo struct {
	createF func(req model.PersonRequest) (int32, error)
	getByIdF func(id int32) (*model.Person, error)
	getAllF func() ([]model.Person, error)
	updateF func(id int32, req model.PersonRequest) (*model.Person, error)
	deleteF func(id int32) error
}

func (m *mockRepo) Create(req model.PersonRequest) (int32, error) {
	return m.createF(req )
}

func (m *mockRepo) GetByID(id int32) (*model.Person, error) {
	return m.getByIdF(id)
}

func (m *mockRepo) GetAll() ([]model.Person, error) {
	return m.getAllF()
}

func (m *mockRepo) Update(id int32, req model.PersonRequest) (*model.Person, error) {
	return m.updateF(id, req)
}

func (m *mockRepo) Delete(id int32) error {
	return m.deleteF(id)
}

func TestPersonService_Create_EmptyName(t *testing.T) {
	svc := NewPersonService(&mockRepo{}) 

	_, err := svc.Create(model.PersonRequest{Name: ""})
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestPersonService_Create_Valid(t *testing.T) {
	rep := &mockRepo{
		createF: func(req model.PersonRequest) (int32, error) {
			if req.Name != "Petr" {
				t.Errorf("expected name Petr, got %s", req.Name)
			}
			return 12, nil
		},
	}

	svc := NewPersonService(rep)

	id, err := svc.Create(model.PersonRequest{Name: "Petr", Age: 18})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 12 {
		t.Errorf("expected id 42, got %d", id)
	}
}

func TestPersonService_GetById_InvalidId (t *testing.T) {
	svc := NewPersonService(&mockRepo{})

	_, err := svc.GetByID(-1) 
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestPersonService_GetById_NotFoundId(t *testing.T) {
	rep := &mockRepo{
		getByIdF: func(id int32) (*model.Person, error) {
			return nil, ErrNotFound
		},
	}
	
	svc:= NewPersonService(rep)

	_, err := svc.GetByID(99)

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestPersonService_Update_NoFields(t *testing.T) {
	rep := &mockRepo{
		updateF: func(id int32, req model.PersonRequest) (*model.Person, error) {
			t.Fatal("repo shouldn`t be called when no fields to update")
			return nil, nil
		},
	}
	svc := NewPersonService(rep)
	_, err := svc.Update(1, model.PersonRequest{})
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}