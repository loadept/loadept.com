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
	ctx context.Context
}

func NewCategoryRepository(rdb *redis.Client, ctx context.Context) *CategoryRepository {
	return &CategoryRepository{
		rdb: rdb,
		ctx: ctx,
	}
}

func (c *CategoryRepository) GetCategories() (*model.CategoryResponse, error) {
	cacheData, err := c.rdb.Get(c.ctx, "categories").Result()
	if err == nil && cacheData != "" {
		var categories model.CategoryResponse
		if err := json.Unmarshal([]byte(cacheData), &categories); err == nil {
			return &categories, nil
		}
	}

	return nil, fmt.Errorf("No results found")
}

func (c *CategoryRepository) StoreCategory(categories model.CategoryResponse) error {
	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		return err
	}
	c.rdb.Set(c.ctx, "categories", categoriesJSON, 0)

	return nil
}
