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
	// newArticle := models.Article{
	// 	Title : article.Title,
	//     Description : article.Description,
	//     Rate : article.Rate,
	// }
	//================================ 1ere methode =============
	// res, err := a.articleRepository.CreateArticle(article)
	// if err != nil {
	// 	return &models.Article{}, errors.New("pas cre")
	// }
	// return &res, nil
	//================================ 2iem methode =============
	return a.articleRepository.CreateArticle(article)
}

// DeleteArticle implements ArticleService
func (a *articleService) DeleteArticle(id string) error {
	//================================ 1ere methode =============
	// err := a.articleRepository.DeleteArticle(id)
	// if err != nil {
	// 	return errors.New("no del")
	// }
	// return nil
	//================================ 2iem methode =============
	return a.articleRepository.DeleteArticle(id)
}

// GetArticle implements ArticleService
func (a *articleService) GetArticle(id string) (*models.Article, error) {
	//================================ 1ere methode =============
	// res, err := a.articleRepository.GetArticle(id)
	// if err != nil {
	// 	return &models.Article{}, errors.New("pad rec")
	// }
	// return &res, nil
	//================================ 2iem methode =============
	return a.articleRepository.GetArticle(id)

}
