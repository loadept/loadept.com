package redis

import (
	"context"
	"encoding/json"
	"fmt"

	domain "github.com/loadept/loadept.com/internal/domain/article"
	"github.com/redis/go-redis/v9"
)

type ArticleRepositoryCache struct {
	rdb *redis.Client
}

func NewArticleRepositoryCache(rdb *redis.Client) domain.ArticleRepositoryCache {
	return &ArticleRepositoryCache{
		rdb: rdb,
	}
}

func (a *ArticleRepositoryCache) GetArticleContent(ctx context.Context, category, articleName string) (string, error) {
	key := fmt.Sprintf("%s:article:%s", category, articleName)

	cacheData, err := a.rdb.Get(ctx, key).Result()
	if err == nil && cacheData != "" {
		return cacheData, nil
	}

	return "", fmt.Errorf("no results found")
}

func (a *ArticleRepositoryCache) GetArticlesByCategorie(ctx context.Context, category string) ([]domain.Article, error) {
	key := fmt.Sprintf("category:%s:articles", category)

	cacheData, err := a.rdb.LRange(ctx, key, 0, -1).Result()
	if err == nil && len(cacheData) > 0 {
		var articles []domain.Article

		for _, articleString := range cacheData {
			var article domain.Article

			if err := json.Unmarshal([]byte(articleString), &article); err == nil {
				articles = append(articles, article)
			}
		}
		return articles, nil
	}

	return nil, fmt.Errorf("no results found")
}

func (a *ArticleRepositoryCache) StoreArticleContent(ctx context.Context, category, articleName, content string) error {
	key := fmt.Sprintf("%s:article:%s", category, articleName)

	if err := a.rdb.Set(ctx, key, content, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (a *ArticleRepositoryCache) StoreArticles(ctx context.Context, category string, articles []domain.Article) error {
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
