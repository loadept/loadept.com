package article

import "context"

type ArticleRepository interface {
	GetArticleContent(ctx context.Context, category, articleName string) (string, error)
	GetArticlesByCategorie(ctx context.Context, category string) ([]Article, error)
}

type ArticleRepositoryCache interface {
	GetArticleContent(ctx context.Context, category, articleName string) (string, error)
	GetArticlesByCategorie(ctx context.Context, category string) ([]Article, error)
	StoreArticleContent(ctx context.Context, category, artcleName, content string) error
	StoreArticles(ctx context.Context, category string, articles []Article) error
}
