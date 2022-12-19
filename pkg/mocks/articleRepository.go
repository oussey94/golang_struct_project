package mocks

import (
	"github.com/oussey94/golang_struct_project/pkg/models"
	"github.com/stretchr/testify/mock"
)

// type ArticleRepository struct {
// 	mock.Mock
// }

// func (m *ArticleRepository) GetArticle(id string) (*models.Article, error) {
// 	ret := m.Called(id)

// 	var r0 models.Article
// 	if rf, ok := ret.Get(0).(func(string) models.Article); ok {
// 		r0 = rf(id)
// 	} else {
// 		r0 = ret.Get(0).(models.Article)
// 	}

// 	var r1 error
// 	if rf, ok := ret.Get(1).(func(string) error); ok {
// 		r1 = rf(id)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return &r0, r1
// }

// func (m *ArticleRepository) CreateArticle(article *models.Article) (*models.Article, error) {
// 	ret := m.Called(article)

// 	var r0 models.Article
// 	if rf, ok := ret.Get(0).(func(*models.Article) models.Article); ok {
// 		r0 = rf(article)
// 	} else {
// 		r0 = ret.Get(0).(models.Article)
// 	}

// 	var r1 error
// 	if rf, ok := ret.Get(1).(func(*models.Article) error); ok {
// 		r1 = rf(article)
// 	} else {
// 		r1 = ret.Error(1)
// 	}

// 	return &r0, r1
// }

// func (m *ArticleRepository) DeleteArticle(id string) error {
// 	ret := m.Called(id)

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(string) error); ok {
// 		r0 = rf(id)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }

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