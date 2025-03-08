package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository"
	"github.com/loadept/loadept.com/pkg/util"
)

type UserService struct {
	repository *repository.UserRepository
	validator  *validator.Validate
}

func NewUserService(repository *repository.UserRepository, validator *validator.Validate) *UserService {
	return &UserService{
		repository: repository,
		validator:  validator,
	}
}

func (s *UserService) RegisterUser(body *model.RegisterUserModel) error {
	err := s.validator.Struct(body)
	if err != nil {
		return util.HandleValidationErrors(err)
	}

	id := uuid.New()
	body.ID = id.String()

	hasPassword, err := util.HasPassword(body.Password)
	if err != nil {
		return err
	}
	body.Password = hasPassword

	err = s.repository.RegisterUser(body)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(body *model.UserModel) (string, error) {
	err := s.validator.Struct(body)
	if err != nil {
		return "", util.HandleValidationErrors(err)
	}

	user, err := s.repository.GetUserByName(body.Username)
	if err != nil {
		return "", err
	}

	if !util.CheckPasswordHash(body.Password, user.Password) {
		return "", fmt.Errorf("Incorrect credentials")
	}

	return user.ID, nil
}
