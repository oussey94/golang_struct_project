package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oussey94/golang_struct_project/pkg/mocks"
	"github.com/oussey94/golang_struct_project/pkg/models"
	"github.com/stretchr/testify/mock"
)

func TestControllerGetById(t *testing.T) {
	// Initialisation du mocker pour le repository
	repo := new(mocks.MockRepository)
	// Initialisation du mocker pour le service
	service := new(mocks.MockService)
	// Initialisation du mocker pour le controller
	ctrl := new(mocks.MockController)

	// Configuration de la réponse attendue pour le repository
	repo.On("GetById", uint(1)).Return(&models.Article{ID: 1, Title: "John", Description: "John", Rate: 5}, nil)
	// Configuration de la réponse attendue pour le service
	service.On("GetById", uint(1)).Return(&models.Article{ID: 1, Title: "John", Description: "John", Rate: 5}, nil)
	// Configuration de la réponse attendue pour la méthode GetById
	ctrl.On("GetById", mock.Anything).Return()

	// Initialisation de la requête HTTP
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/v1/articles/1", nil)

	// Création d'un contexte gin à partir de la requête HTTP
	c, _ := gin.CreateTestContext(w)
	c.Request = r
	// Exécution de la méthode du controller
	ctrl.GetById(c)

	// Vérification que la méthode a bien été appelée
	//ctrl.AssertNumberOfCalls(t, "GetById", 1)
	// Exécution de la méthode du controller
	//ctrl.GetById(w, r)

}
