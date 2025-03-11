package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/redis/go-redis/v9"
)

type ArticleService struct {
	rdb         *redis.Client
	ctx         context.Context
	httpClient  *http.Client
	baseURL     string
	githubURL   string
	githubToken string
}

func NewArticleService(httpClient *http.Client, rdb *redis.Client, ctx context.Context) *ArticleService {
	return &ArticleService{
		rdb:         rdb,
		ctx:         ctx,
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubURL:   config.Env.GITHUB_URL,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *ArticleService) GetArticleByName(category, name string) (io.ReadCloser, error) {
	endPoint := fmt.Sprintf("contents/articles/%s/%s.md", category, name)
	fullURL := fmt.Sprintf("%s/%s", s.baseURL, endPoint)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Accept", "application/vnd.github.raw+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Error to request api: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (s *ArticleService) GetArticles(category string) ([]model.ArticleModel, error) {
	key := fmt.Sprintf("category:%s:articles", category)
	{ // Get from cache if exists
		cacheData, err := s.rdb.LRange(s.ctx, key, 0, -1).Result()
		if err == nil && len(cacheData) > 0 {
			var articles []model.ArticleModel
			for _, articleString := range cacheData {
				var article model.ArticleModel
				if err := json.Unmarshal([]byte(articleString), &article); err == nil {
					articles = append(articles, article)
				}
			}
			return articles, nil
		}
	}

	endPoint := fmt.Sprintf("articles/%s", category)
	fullURL := fmt.Sprintf("%s/contents/%s", s.baseURL, url.PathEscape(endPoint))
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Accept", "application/vnd.github.raw+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error to request api: %d", resp.StatusCode)
	}

	var files []struct {
		Name string `json:"name"`
		Path string `json:"path"`
		SHA  string `json:"sha"`
		Type string `json:"type"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, err
	}

	articles := []model.ArticleModel{}
	var commitRequests []*http.Request
	for _, file := range files {
		if file.Type != "file" || !strings.HasSuffix(file.Name, ".md") {
			continue
		}

		article := model.ArticleModel{
			Sha:  file.SHA,
			Name: file.Name,
			Path: file.Path,
		}

		fullCommitURL := fmt.Sprintf("%s/commits?path=%s&page=1&per_page=1", s.baseURL, url.PathEscape(file.Path))
		commitReq, err := http.NewRequest("GET", fullCommitURL, nil)
		if err != nil {
			return nil, err
		}

		commitReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
		commitReq.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		commitReq.Header.Set("Accept", "application/vnd.github.raw+json")

		commitRequests = append(commitRequests, commitReq)
		articles = append(articles, article)
	}

	for i, commitReq := range commitRequests {
		commitResp, err := s.httpClient.Do(commitReq)
		if err != nil {
			continue
		}
		defer commitResp.Body.Close()

		if commitResp.StatusCode != http.StatusOK {
			continue
		}

		var commits []struct {
			Commit struct {
				Committer struct {
					Date time.Time `json:"date"`
				} `json:"committer"`
			} `json:"commit"`
		}

		if err := json.NewDecoder(commitResp.Body).Decode(&commits); err == nil && len(commits) > 0 {
			articles[i].UpdatedAt = commits[0].Commit.Committer.Date
		}
	}
	{ // Store articles in cache
		pipe := s.rdb.Pipeline()
		for _, article := range articles {
			articleJSON, err := json.Marshal(article)
			if err == nil {
				pipe.RPush(s.ctx, key, articleJSON)
			}
		}
		pipe.Exec(s.ctx)
	}

	return articles, nil
}

func (s *ArticleService) EditArticle(category, name string) string {
	endPoint := fmt.Sprintf("%s/%s", category, name)
	endPoint = url.PathEscape(endPoint)

	editURL := fmt.Sprintf("%s/blob/main/articles/%s.md", s.githubURL, endPoint)
	return editURL
}
