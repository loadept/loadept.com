package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/loadept/loadept.com/api/v1/handler"
	application "github.com/loadept/loadept.com/internal/application/article"
	"github.com/loadept/loadept.com/internal/config"
	domain "github.com/loadept/loadept.com/internal/domain/article"
	"github.com/loadept/loadept.com/internal/infrastructure/cache"
	infraExternal "github.com/loadept/loadept.com/internal/infrastructure/repository/external"
	infraRedis "github.com/loadept/loadept.com/internal/infrastructure/repository/redis"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRawArticleByName(t *testing.T) {
	err := os.MkdirAll("logs", os.ModePerm)
	assert.NoError(t, err, "Should be able to create the directory records")
	defer os.RemoveAll("logs")

	config.LoadEnviron()
	logger.NewLogger()

	httpClient := &http.Client{}

	ctx := context.Background()
	redisClient, err := cache.NewRedisConnection(ctx)
	if err != nil {
		t.Errorf("Error to connect with redis: %v", err)
	}

	rdb := redisClient.GetClient()

	err = rdb.Set(ctx, "test:article:test-article", "hello test", time.Second*30).Err()
	require.NoError(t, err)

	repositoryExternal := infraExternal.NewArticleRepository(httpClient)
	repositoryRedis := infraRedis.NewArticleRepositoryCache(rdb)
	service := application.NewArticleService(repositoryExternal, repositoryRedis)
	handler := handler.NewApiArticlesHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /articles/{category}/{name}", handler.GetRawArticleByName)

	t.Run("Get Article", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Get(server.URL + "/articles/test/test-article")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body), 1)
	})

	t.Run("HEAD Method", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Head(server.URL + "/articles/test/test-article")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Empty(t, bodyBytes)
	})

	t.Run("Get Article with Wrong Method", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Post(server.URL+"/articles/test/test-article", "application/json", nil)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, body, []byte("Method Not Allowed"))
	})

	t.Run("Get Non-Existent Article", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Get(server.URL + "/articles/test/non-existent")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)

		var body struct {
			Detail string `json:"detail"`
		}
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Detail), 1)
	})
}

func TestGetListArticle(t *testing.T) {
	config.LoadEnviron()

	httpClient := &http.Client{}

	ctx := context.Background()
	redisClient, err := cache.NewRedisConnection(ctx)
	if err != nil {
		t.Errorf("Error to connect with redis: %v", err)
	}

	rdb := redisClient.GetClient()

	testData := domain.Article{
		Name:    "test",
		Path:    "test/",
		Sha:     "sha123321",
		HtmlURL: "http://test.com/test.html",
	}
	dataByte, err := json.Marshal(testData)
	if err != nil {
		t.Errorf("Error to marshal json: %v", err)
	}

	err = rdb.RPush(ctx, "category:test:articles", dataByte).Err()
	require.NoError(t, err)
	err = rdb.Expire(ctx, "category:test:articles", time.Second*30).Err()
	require.NoError(t, err)

	repositoryExternal := infraExternal.NewArticleRepository(httpClient)
	repositoryRedis := infraRedis.NewArticleRepositoryCache(rdb)
	service := application.NewArticleService(repositoryExternal, repositoryRedis)
	handler := handler.NewApiArticlesHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /articles/{category}", handler.GetListArticles)

	t.Run("Get List Article", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Get(server.URL + "/articles/test")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body struct {
			Articles []domain.Article `json:"articles"`
		}
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Articles), 1)
	})

	t.Run("HEAD Method", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Head(server.URL + "/articles/test")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Empty(t, bodyBytes)
	})

	t.Run("Listing with Wrong Method", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Post(server.URL+"/articles/test", "application/json", nil)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, body, []byte("Method Not Allowed"))
	})

	t.Run("Listing Non-Existent Articles", func(t *testing.T) {
		server := httptest.NewServer(mux)
		defer server.Close()

		resp, err := http.Get(server.URL + "/articles/non-existent-category")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)

		var body struct {
			Detail string `json:"detail"`
		}
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Detail), 1)
	})
}
