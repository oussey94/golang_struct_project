package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oussey94/golang_struct_project/pkg/controllers"
	"github.com/oussey94/golang_struct_project/pkg/db"
	"github.com/oussey94/golang_struct_project/pkg/repository"
	"github.com/oussey94/golang_struct_project/pkg/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	articleRepository := repository.NewArticleRepository(db.DB)
	articleService := services.NewArticleService(articleRepository)
	articleController := controllers.NewArticleControlller(articleService)

	r.GET("/api/v1/articles/:id", articleController.GetArticle)
	r.POST("/api/v1/articles", articleController.PostArticle)
	r.DELETE("/api/v1/articles/:id", articleController.DeleteArticle)

	return r
}
