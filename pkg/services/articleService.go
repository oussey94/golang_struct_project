package services

import (
	"github.com/oussey94/golang_struct_project/pkg/models"
	"github.com/oussey94/golang_struct_project/pkg/repository"
)

type ArticleService interface {
	//GetArticles() ([]models.Article, bool, error)
	GetArticle(id string) (*models.Article, error)
	CreateArticle(article *models.Article) (*models.Article, error)
	DeleteArticle(id string) error
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleService{articleRepository: articleRepo}
}

// CreateArticle implements ArticleService
func (a *articleService) CreateArticle(article *models.Article) (*models.Article, error) {

	return a.articleRepository.CreateArticle(article)
}

// DeleteArticle implements ArticleService
func (a *articleService) DeleteArticle(id string) error {

	return a.articleRepository.DeleteArticle(id)
}

// GetArticle implements ArticleService
func (a *articleService) GetArticle(id string) (*models.Article, error) {

	return a.articleRepository.GetArticle(id)

}
