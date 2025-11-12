package sitemap

import (
	"context"
	"net/url"

	articleDomain "github.com/loadept/loadept.com/internal/domain/article"
	categoryDomain "github.com/loadept/loadept.com/internal/domain/category"
)

type SitemapService struct {
	categorieRepository categoryDomain.CategoryRepository
	articleRepository   articleDomain.ArticleRepository
}

func NewSitemapService(categorieRepository categoryDomain.CategoryRepository, articleRepository articleDomain.ArticleRepository) *SitemapService {
	return &SitemapService{
		categorieRepository: categorieRepository,
		articleRepository:   articleRepository,
	}
}

func (s *SitemapService) GetArticles(ctx context.Context) ([]string, error) {
	categories, err := s.categorieRepository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	var articles []string
	for _, category := range categories.Category {
		articleRepository, err := s.articleRepository.GetArticlesByCategorie(ctx, category.Name)
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
