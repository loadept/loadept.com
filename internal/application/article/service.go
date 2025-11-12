package article

import (
	"context"
	"io"
	"strings"

	domain "github.com/loadept/loadept.com/internal/domain/article"
)

type ArticleService struct {
	repository      domain.ArticleRepository
	repositoryCache domain.ArticleRepositoryCache
}

func NewArticleService(repository domain.ArticleRepository, repositoryCache domain.ArticleRepositoryCache) *ArticleService {
	return &ArticleService{
		repository:      repository,
		repositoryCache: repositoryCache,
	}
}

func (s *ArticleService) GetArticleContent(ctx context.Context, category, artcleName string) (io.ReadCloser, error) {
	cacheContent, err := s.repositoryCache.GetArticleContent(ctx, category, artcleName)
	if err == nil && len(cacheContent) > 0 {
		reader := strings.NewReader(cacheContent)
		return io.NopCloser(reader), nil
	}

	content, err := s.repository.GetArticleContent(ctx, category, artcleName)
	if err != nil {
		return nil, err
	}

	if err := s.repositoryCache.StoreArticleContent(ctx, category, artcleName, content); err != nil {
		return nil, err
	}

	reader := strings.NewReader(content)
	return io.NopCloser(reader), nil
}

func (s *ArticleService) GetArticlesByCategorie(ctx context.Context, category string) ([]domain.Article, error) {
	cacheArticles, err := s.repositoryCache.GetArticlesByCategorie(ctx, category)
	if err == nil && len(cacheArticles) > 0 {
		return cacheArticles, nil
	}

	articles, err := s.repository.GetArticlesByCategorie(ctx, category)
	if err != nil {
		return nil, err
	}

	if err := s.repositoryCache.StoreArticles(ctx, category, articles); err != nil {
		return nil, err
	}

	return articles, nil
}
