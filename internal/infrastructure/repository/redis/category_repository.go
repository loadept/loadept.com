package redis

import (
	"context"
	"encoding/json"
	"fmt"

	domain "github.com/loadept/loadept.com/internal/domain/category"
	"github.com/redis/go-redis/v9"
)

type CategoryRepository struct {
	rdb *redis.Client
}

func NewCategoryRepository(rdb *redis.Client) domain.CategoryRepositoryCache {
	return &CategoryRepository{
		rdb: rdb,
	}
}

func (c *CategoryRepository) GetCategories(ctx context.Context) (*domain.CategoryList, error) {
	cacheData, err := c.rdb.Get(ctx, "categories").Result()
	if err == nil && cacheData != "" {
		var categories domain.CategoryList
		if err := json.Unmarshal([]byte(cacheData), &categories); err == nil {
			return &categories, nil
		}
	}

	return nil, fmt.Errorf("no results found")
}

func (c *CategoryRepository) StoreCategory(ctx context.Context, categories *domain.CategoryList) error {
	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		return err
	}
	c.rdb.Set(ctx, "categories", categoriesJSON, 0)

	return nil
}
