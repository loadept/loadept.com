package di

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/redis/go-redis/v9"

	"github.com/loadept/loadept.com/api/v1/handler"
	articleApp "github.com/loadept/loadept.com/internal/application/article"
	categoryApp "github.com/loadept/loadept.com/internal/application/category"
	checkHealthApp "github.com/loadept/loadept.com/internal/application/checkhealth"
	sitemapApp "github.com/loadept/loadept.com/internal/application/sitemap"
	articleDomain "github.com/loadept/loadept.com/internal/domain/article"
	categoryDomain "github.com/loadept/loadept.com/internal/domain/category"
	checkHealthDomain "github.com/loadept/loadept.com/internal/domain/checkhealth"
	infraDB "github.com/loadept/loadept.com/internal/infrastructure/repository/db"
	infra "github.com/loadept/loadept.com/internal/infrastructure/repository/external"
	infraCache "github.com/loadept/loadept.com/internal/infrastructure/repository/redis"
)

type Container struct {
	DB         *sql.DB
	Redis      *redis.Client
	HTTPClient *http.Client
	Proxy      *httputil.ReverseProxy

	// Repositories
	articleRepo          articleDomain.ArticleRepository
	articleRepoCache     articleDomain.ArticleRepositoryCache
	categoryRepo         categoryDomain.CategoryRepository
	categoryRepoCache    categoryDomain.CategoryRepositoryCache
	checkHealthRepo      checkHealthDomain.CheckHealthRepository
	checkHealthRepoCache checkHealthDomain.CheckHealthRepository

	// Services
	articleService     *articleApp.ArticleService
	categoryService    *categoryApp.CategoryService
	checkHealthService *checkHealthApp.CheckHealthService
	sitemapService     *sitemapApp.SitemapService

	// Handlers
	ArticleHandler  *handler.ApiArticleHandler
	CategoryHandler *handler.ApiCategoriesHandler
	HealthHandler   *handler.CheckHealthHandler
	SitemapHandler  *handler.SitemapHandler
	ApiPDFHandler   *handler.ApiPDFHandler

	once sync.Once
}

func NewContainer(
	db *sql.DB,
	redis *redis.Client,
	httpClient *http.Client,
	proxy *httputil.ReverseProxy,
) *Container {
	return &Container{
		DB:         db,
		Redis:      redis,
		HTTPClient: httpClient,
		Proxy:      proxy,
	}
}

func (c *Container) Build(ctx context.Context) {
	c.once.Do(func() {
		c.buildDependencies()
	})
}

func (c *Container) buildDependencies() {
	c.articleRepo = infra.NewArticleRepository(c.HTTPClient)
	c.articleRepoCache = infraCache.NewArticleRepositoryCache(c.Redis)
	c.categoryRepo = infra.NewCategoryRepository(c.HTTPClient)
	c.categoryRepoCache = infraCache.NewCategoryRepository(c.Redis)
	c.checkHealthRepo = infraDB.NewCheckHealthDBRepository(c.DB)
	c.checkHealthRepoCache = infraCache.NewCheckHealthRedisRepository(c.Redis)

	c.articleService = articleApp.NewArticleService(
		c.articleRepo,
		c.articleRepoCache,
	)

	c.categoryService = categoryApp.NewCategoryService(
		c.categoryRepo,
		c.categoryRepoCache,
	)

	c.checkHealthService = checkHealthApp.NewCheckHealthService(
		c.checkHealthRepo,
		c.checkHealthRepoCache,
	)

	c.sitemapService = sitemapApp.NewSitemapService(
		c.categoryRepo,
		c.articleRepo,
	)

	c.ArticleHandler = handler.NewApiArticlesHandler(c.articleService)
	c.CategoryHandler = handler.NewApiCategoryHandler(c.categoryService)
	c.HealthHandler = handler.NewHealthHandler(c.checkHealthService)
	c.SitemapHandler = handler.NewSitemapHandler(c.sitemapService)
	c.ApiPDFHandler = handler.NewApiPDFHandler(c.Proxy)
}
