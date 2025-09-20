package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/redis/go-redis/v9"
)

type ArticleRepository struct {
	rdb *redis.Client
}

func NewArticleRepository(rdb *redis.Client) *ArticleRepository {
	return &ArticleRepository{
		rdb: rdb,
	}
}

func (a *ArticleRepository) GetRawArticle(ctx context.Context, articleName string) (string, error) {
	key := fmt.Sprintf("article:%s", articleName)

	cacheData, err := a.rdb.Get(ctx, key).Result()
	if err == nil && cacheData != "" {
		return cacheData, nil
	}

	return "", fmt.Errorf("No results found")
}

func (a *ArticleRepository) StoreRawArticle(ctx context.Context, articleName string, data string) error {
	key := fmt.Sprintf("article:%s", articleName)

	if err := a.rdb.Set(ctx, key, data, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepository) GetListArticleByCategory(ctx context.Context, category string) ([]model.ArticleModel, error) {
	key := fmt.Sprintf("category:%s:articles", category)

	cacheData, err := a.rdb.LRange(ctx, key, 0, -1).Result()
	if err == nil && len(cacheData) > 0 {
		var articles []model.ArticleModel

		for _, articleString := range cacheData {
			var article model.ArticleModel

			if err := json.Unmarshal([]byte(articleString), &article); err == nil {
				articles = append(articles, article)
			}
		}
		return articles, nil
	}

	return nil, fmt.Errorf("No results found")
}

func (a *ArticleRepository) StoreListArticles(ctx context.Context, category string, articles []model.ArticleModel) error {
	key := fmt.Sprintf("category:%s:articles", category)

	pipe := a.rdb.Pipeline()
	for _, article := range articles {
		articleJSON, err := json.Marshal(article)
		if err != nil {
			return err
		}
		pipe.RPush(ctx, key, articleJSON)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
