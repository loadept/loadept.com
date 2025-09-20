package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loadept/loadept.com/api/handler"
	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/infrastructure/cache"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository/redis"
	"github.com/loadept/loadept.com/internal/service"
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
	repository := redis.NewCategoryRepository(rdb)
	service := service.NewCategoryService(httpClient, repository)
	handler := handler.NewApiCategoryHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/categories", handler.GetCategories)

	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("Get Categories", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/categories")
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var body model.CategoryResponse
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Category), 1)
	})

	t.Run("Wrong Method", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/categories", "application/json", nil)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

		var body struct {
			Detail string `json:"detail"`
		}
		err = json.NewDecoder(resp.Body).Decode(&body)
		require.NoError(t, err)

		assert.GreaterOrEqual(t, len(body.Detail), 1)
	})
}
