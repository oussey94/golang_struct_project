package repository

import (
	"errors"

	"github.com/oussey94/golang_struct_project/pkg/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	//GetArticles() ([]models.Article, bool, error)
	GetArticle(id string) (*models.Article, error)
	CreateArticle(article *models.Article) (*models.Article, error)
	DeleteArticle(id string) error
}

type articleRepository struct {
	DB *gorm.DB
}

//get injected database
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{DB: db}
}

func (a *articleRepository) GetArticle(id string) (*models.Article, error) {
	var article models.Article
	res := a.DB.First(&article, id)
	if res.RowsAffected == 0 {
		return &models.Article{}, errors.New("article not found")
	}
	return &article, nil
}

func (a *articleRepository) CreateArticle(article *models.Article) (*models.Article, error) {
	res := a.DB.Create(article)
	if res.RowsAffected == 0 {
		return &models.Article{}, errors.New("article not created")
	}
	return article, nil
}

func (a *articleRepository) DeleteArticle(id string) error {
	var article *models.Article
	res := a.DB.Where(id).Delete(&article)
	if res.RowsAffected == 0 {
		return errors.New("article data not deleted")
	}
	return nil
}
