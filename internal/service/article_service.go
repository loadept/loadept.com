package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/model"
)

type ArticleService struct {
	httpClient  *http.Client
	baseURL     string
	githubURL   string
	githubToken string
}

func NewArticleService(httpClient *http.Client) *ArticleService {
	return &ArticleService{
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubURL:   config.Env.GITHUB_URL,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *ArticleService) GetArticleBySha(category, name string) (io.ReadCloser, error) {
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
	if len(category) == 0 {
		return []model.ArticleModel{}, nil
	}

	endPoint := fmt.Sprintf("articles/%s", category)
	endPoint = url.PathEscape(endPoint)

	fullURL := fmt.Sprintf("%s/contents/%s", s.baseURL, endPoint)

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
	for _, file := range files {
		if file.Type != "file" || !strings.HasSuffix(file.Name, ".md") {
			continue
		}

		article := model.ArticleModel{
			Sha:   file.SHA,
			Name: file.Name,
			Path: file.Path,
		}

		commitEndPoint := fmt.Sprintf("articles/%s/%s", category, file.Name)
		commitEndPoint = url.PathEscape(commitEndPoint)

		fullCommitURL := fmt.Sprintf("%s/commits?path=%s&page=1&per_page=1", s.baseURL, commitEndPoint)

		commitReq, err := http.NewRequest("GET", fullCommitURL, nil)
		if err != nil {
			return nil, err
		}

		commitReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
		commitReq.Header.Set("X-GitHub-Api-Version", "2022-11-28")
		commitReq.Header.Set("Accept", "application/vnd.github.raw+json")

		commitResp, err := s.httpClient.Do(commitReq)
		if err != nil {
			return nil, err
		}

		var commits []struct {
			Commit struct {
				Committer struct {
					Date time.Time `json:"date"`
				} `json:"committer"`
			} `json:"commit"`
		}

		if err := json.NewDecoder(commitResp.Body).Decode(&commits); err != nil {
			commitResp.Body.Close()
			return nil, err
		}
		commitResp.Body.Close()

		if len(commits) > 0 {
			article.UpdatedAt = commits[0].Commit.Committer.Date
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (s *ArticleService) EditArticle(category, name string) string {
	endPoint := fmt.Sprintf("%s/%s", category, name)
	endPoint = url.PathEscape(endPoint)

	editURL := fmt.Sprintf("%s/blob/main/articles/%s.md", s.githubURL, endPoint)
	return editURL
}
