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
	ctx context.Context
}

func NewArticleRepository(rdb *redis.Client, ctx context.Context) *ArticleRepository {
	return &ArticleRepository{
		rdb: rdb,
		ctx: ctx,
	}
}

func (a *ArticleRepository) GetRawArticle(articleName string) (string, error) {
	key := fmt.Sprintf("article:%s", articleName)

	cacheData, err := a.rdb.Get(a.ctx, key).Result()
	if err == nil && cacheData != "" {
		return cacheData, nil
	}

	return "", fmt.Errorf("No results found")
}

func (a *ArticleRepository) StoreRawArticle(articleName string, data string) error {
	key := fmt.Sprintf("article:%s", articleName)

	if err := a.rdb.Set(a.ctx, key, data, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepository) GetListArticleByCategory(category string) ([]model.ArticleModel, error) {
	key := fmt.Sprintf("category:%s:articles", category)

	cacheData, err := a.rdb.LRange(a.ctx, key, 0, -1).Result()
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

func (a *ArticleRepository) StoreListArticles(category string, articles []model.ArticleModel) error {
	key := fmt.Sprintf("category:%s:articles", category)

	pipe := a.rdb.Pipeline()
	for _, article := range articles {
		articleJSON, err := json.Marshal(article)
		if err != nil {
			return err
		}
		pipe.RPush(a.ctx, key, articleJSON)
	}
	_, err := pipe.Exec(a.ctx)
	if err != nil {
		return err
	}

	return nil
}
