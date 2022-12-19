package mocks

import (
	"github.com/oussey94/golang_struct_project/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll() ([]*models.Article, error) {
	args := m.Called()
	return args.Get(0).([]*models.Article), args.Error(1)
}

func (m *MockRepository) GetById(id int) (*models.Article, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Article), args.Error(1)
}

func (m *MockRepository) Create(p *models.Article) (*models.Article, error) {
	args := m.Called(p)
	return args.Get(0).(*models.Article), args.Error(1)
}

func (m *MockRepository) Update(p *models.Article) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
