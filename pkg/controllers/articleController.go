package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oussey94/golang_struct_project/pkg/models"
	"github.com/oussey94/golang_struct_project/pkg/services"
)

type ArticleController interface {
	PostArticle(c *gin.Context)
	GetArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
}
type articleController struct {
	articleService services.ArticleService
}

func NewArticleControlller(articleContr services.ArticleService) ArticleController {
	return &articleController{articleService: articleContr}
}

func (ac *articleController) PostArticle(c *gin.Context) {
	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err,})
		return
	}
	res, err := ac.articleService.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err,})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": res,})
}

func (ac *articleController) GetArticle(c *gin.Context) {
	id := c.Param("id")
	res, err := ac.articleService.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found",})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": res,})
}

func (ac *articleController) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := ac.articleService.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found",})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "article deleted successfully",})
}