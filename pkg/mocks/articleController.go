package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockController struct {
	mock.Mock
}

func (m *MockController) GetAll(c *gin.Context) {
	m.Called(c)
}
func (m *MockController) GetById(c *gin.Context) {
	m.Called(c)
}
func (m *MockController) Create(c *gin.Context) {
	m.Called(c)
}
func (m *MockController) Update(c *gin.Context) {
	m.Called(c)
}
func (m *MockController) Delete(c *gin.Context) {
	m.Called(c)
}
