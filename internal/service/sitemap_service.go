package service

import (
	"context"
	"net/url"
)

type SitemapService struct {
	categorieService *CategoryService
	articleService   *ArticleService
}

func NewSitemapService(categorieService *CategoryService, articleService *ArticleService) *SitemapService {
	return &SitemapService{
		categorieService: categorieService,
		articleService:   articleService,
	}
}

func (s *SitemapService) GetArticles(ctx context.Context) ([]string, error) {
	categoryService, err := s.categorieService.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	var articles []string
	for _, category := range categoryService.Category {
		articleRepository, err := s.articleService.GetListArticles(ctx, category.Name)
		if err != nil {
			return nil, err
		}

		for _, article := range articleRepository {
			urlName := url.URL{Path: article.Path}
			articles = append(articles, urlName.String())
		}
	}

	return articles, nil
}
