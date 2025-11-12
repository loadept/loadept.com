package category

import (
	"context"

	domain "github.com/loadept/loadept.com/internal/domain/category"
)

type CategoryService struct {
	repository      domain.CategoryRepository
	repositoryCache domain.CategoryRepositoryCache
}

func NewCategoryService(repository domain.CategoryRepository, repositoryCache domain.CategoryRepositoryCache) *CategoryService {
	return &CategoryService{
		repository:      repository,
		repositoryCache: repositoryCache,
	}
}

func (s *CategoryService) GetCategories(ctx context.Context) (*domain.CategoryList, error) {
	cacheCategory, err := s.repositoryCache.GetCategories(ctx)
	if err == nil && len(cacheCategory.Category) > 0 {
		return cacheCategory, nil
	}

	categoryList, err := s.repository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	if err := s.repositoryCache.StoreCategory(ctx, categoryList); err != nil {
		return nil, err
	}

	return categoryList, nil
}
