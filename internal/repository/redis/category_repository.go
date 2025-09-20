package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/redis/go-redis/v9"
)

type CategoryRepository struct {
	rdb *redis.Client
}

func NewCategoryRepository(rdb *redis.Client) *CategoryRepository {
	return &CategoryRepository{
		rdb: rdb,
	}
}

func (c *CategoryRepository) GetCategories(ctx context.Context) (*model.CategoryResponse, error) {
	cacheData, err := c.rdb.Get(ctx, "categories").Result()
	if err == nil && cacheData != "" {
		var categories model.CategoryResponse
		if err := json.Unmarshal([]byte(cacheData), &categories); err == nil {
			return &categories, nil
		}
	}

	return nil, fmt.Errorf("no results found")
}

func (c *CategoryRepository) StoreCategory(ctx context.Context, categories model.CategoryResponse) error {
	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		return err
	}
	c.rdb.Set(ctx, "categories", categoriesJSON, 0)

	return nil
}
