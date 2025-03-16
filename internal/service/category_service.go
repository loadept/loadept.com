package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository/redis"
)

type CategoryService struct {
	respository *redis.CategoryRepository
	httpClient  *http.Client
	baseURL     string
	githubToken string
}

func NewCategoryService(httpClient *http.Client, repository *redis.CategoryRepository) *CategoryService {
	return &CategoryService{
		respository: repository,
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *CategoryService) GetCategories() (*model.CategoryResponse, error) {
	cacheCategory, err := s.respository.GetCategories()
	if err == nil {
		return cacheCategory, nil
	}

	fullURL := fmt.Sprintf("%s/contents/metadata.json", s.baseURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Accept", "application/vnd.github.raw+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error to request api: %d", resp.StatusCode)
	}

	var categories model.CategoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	if err := s.respository.StoreCategory(categories); err != nil {
		return nil, err
	}

	return &categories, nil
}
