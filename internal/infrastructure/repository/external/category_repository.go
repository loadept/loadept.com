package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/loadept/loadept.com/internal/config"
	domain "github.com/loadept/loadept.com/internal/domain/category"
)

type CategoryRepository struct {
	httpClient  *http.Client
	baseURL     string
	githubToken string
}

func NewCategoryRepository(httpClient *http.Client) domain.CategoryRepository {
	return &CategoryRepository{
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *CategoryRepository) GetCategories(ctx context.Context) (*domain.CategoryList, error) {
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
		return nil, fmt.Errorf("error to request api: %d", resp.StatusCode)
	}

	var categories domain.CategoryList
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return &categories, nil
}
