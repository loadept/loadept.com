package service

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository"
	"github.com/loadept/loadept.com/pkg/util"
)

type CategoryService struct {
	repository *repository.CategoryRepository
	validator  *validator.Validate
}

func NewCategoryService(repository *repository.CategoryRepository, validator *validator.Validate) *CategoryService {
	return &CategoryService{
		repository: repository,
		validator:  validator,
	}
}

func (s *CategoryService) RegisterCategory(body *model.CategoryModel) error {
	err := s.validator.Struct(body)
	if err != nil {
		return util.HandleValidationErrors(err)
	}

	id := uuid.New()
	body.ID = id.String()

	err = s.repository.RegisterCategory(body)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetCategoryByID(id string) (*model.CategoryModel, error) {
	category, err := s.repository.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategories(page string) ([]*model.CategoryModel, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	categories, err := s.repository.SelectCategories(6, pageInt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return categories, nil
}
