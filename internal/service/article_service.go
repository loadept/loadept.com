package service

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository"
	"github.com/loadept/loadept.com/pkg/util"
)

type ArticleService struct {
	repository *repository.ArticleRepository
	validator  *validator.Validate
}

func NewArticleService(repository *repository.ArticleRepository, validator *validator.Validate) *ArticleService {
	return &ArticleService{
		repository: repository,
		validator:  validator,
	}
}

func (s *ArticleService) RegisterArticle(userID string, body *model.ArticleModel) error {
	err := s.validator.Struct(body)
	if err != nil {
		return util.HandleValidationErrors(err)
	}

	id := uuid.New()
	body.ID = id.String()

	err = s.repository.RegisterArticle(userID, body)
	if err != nil {
		return err
	}

	return nil
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
