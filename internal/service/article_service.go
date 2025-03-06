package service

import (
	"strconv"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository"
)

type ArticleService struct {
	repository *repository.ArticleRepository
}

func NewArticleService(repository *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		repository: repository,
	}
}

func (s *ArticleService) RegisterArticle() {

}

func (s *ArticleService) GetArticleByID(articleID string) (*model.ArticleModel, error) {
	article, err := s.repository.GetArticleByID(articleID)
	if err != nil {
		return nil, err
	}

	return article, err
}

func (s *ArticleService) GetArticles(categoryName, articleTitle string, page string) ([]*model.ArticleModel, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	articles, err := s.repository.SelectArticles(categoryName, articleTitle, 10, pageInt)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *ArticleService) GetRecentArticles(categoryName string) ([]*model.ArticleModel, error) {
	articles, err := s.repository.SelectArticles(categoryName, "", 5, 1)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
