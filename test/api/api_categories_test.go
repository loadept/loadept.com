package test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loadept/loadept.com/api/v1/handler"
	application "github.com/loadept/loadept.com/internal/application/category"
	"github.com/loadept/loadept.com/internal/config"
	domain "github.com/loadept/loadept.com/internal/domain/category"
	"github.com/loadept/loadept.com/internal/infrastructure/cache"
	infraExternal "github.com/loadept/loadept.com/internal/infrastructure/repository/external"
	infraRedis "github.com/loadept/loadept.com/internal/infrastructure/repository/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	config.LoadEnviron()

	httpClient := &http.Client{}

	ctx := context.Background()
	redisClient, err := cache.NewRedisConnection(ctx)
	if err != nil {
		t.Errorf("Error to connect with redis: %v", err)
	}

	rdb := redisClient.GetClient()
	repositoryExternal := infraExternal.NewCategoryRepository(httpClient)
	repositoryRedis := infraRedis.NewCategoryRepository(rdb)
	service := application.NewCategoryService(repositoryExternal, repositoryRedis)
	handler := handler.NewApiCategoryHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /categories", handler.GetCategories)

	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("Get Categories", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/categories")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body domain.CategoryList
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Category), 1)
	})

	t.Run("Wrong Method", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/categories", "application/json", nil)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, body, []byte("Method Not Allowed"))
	})
}
