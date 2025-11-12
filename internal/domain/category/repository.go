package category

import "context"

type CategoryRepository interface {
	GetCategories(ctx context.Context) (*CategoryList, error)
}

type CategoryRepositoryCache interface {
	GetCategories(ctx context.Context) (*CategoryList, error)
	StoreCategory(ctx context.Context, categories *CategoryList) error
}
