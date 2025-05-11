package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/repository/redis"
)

type ArticleService struct {
	repository  *redis.ArticleRepository
	httpClient  *http.Client
	baseURL     string
	githubToken string
}

func NewArticleService(httpClient *http.Client, repository *redis.ArticleRepository) *ArticleService {
	return &ArticleService{
		repository:  repository,
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *ArticleService) GetArticleByName(category, name string) (io.ReadCloser, error) {
	cacheArticle, err := s.repository.GetArticle(name)
	if err == nil && len(cacheArticle) > 0 {
		reader := strings.NewReader(cacheArticle)
		return io.NopCloser(reader), nil
	}

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

	reader := bufio.NewReader(resp.Body)
	var builder strings.Builder
	buffer := make([]byte, 4096)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		builder.Write(buffer[:n])
	}
	if err := s.repository.StoreArticle(name, builder.String()); err != nil {
		return nil, err
	}

	newBody := bytes.NewReader(buffer)
	return io.NopCloser(newBody), nil
}

func (s *ArticleService) GetArticles(category string) ([]model.ArticleModel, error) {
	cacheArticles, err := s.repository.GetArticleByCategory(category)
	if err == nil && len(cacheArticles) > 0 {
		return cacheArticles, nil
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
		Name    string `json:"name"`
		Path    string `json:"path"`
		SHA     string `json:"sha"`
		HtmlURL string `json:"html_url,omitempty"`
		Type    string `json:"type"`
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
			Sha:     file.SHA,
			Name:    file.Name,
			Path:    file.Path,
			HtmlURL: file.HtmlURL,
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

	if err := s.repository.StoreArticles(category, articles); err != nil {
		return nil, err
	}

	return articles, nil
}
