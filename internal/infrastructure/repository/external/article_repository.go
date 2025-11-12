package external

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	domain "github.com/loadept/loadept.com/internal/domain/article"
)

type ArticleRepository struct {
	httpClient  *http.Client
	baseURL     string
	githubToken string
}

func NewArticleRepository(httpClient *http.Client) domain.ArticleRepository {
	return &ArticleRepository{
		httpClient:  httpClient,
		baseURL:     config.Env.GITHUB_API,
		githubToken: config.Env.GITHUB_TOKEN,
	}
}

func (s *ArticleRepository) GetArticleContent(ctx context.Context, category, articleName string) (string, error) {
	endPoint := fmt.Sprintf("contents/articles/%s/%s.md", category, articleName)
	fullURL := fmt.Sprintf("%s/%s", s.baseURL, endPoint)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.githubToken))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Accept", "application/vnd.github.raw+json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "", fmt.Errorf("error to request api: %d", resp.StatusCode)
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

	return builder.String(), nil
}

func (s *ArticleRepository) GetArticlesByCategorie(ctx context.Context, category string) ([]domain.Article, error) {
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
		return nil, fmt.Errorf("error to request api: %d", resp.StatusCode)
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

	var articles []domain.Article
	var commitRequests []*http.Request
	for _, file := range files {
		if file.Type != "file" || !strings.HasSuffix(file.Name, ".md") {
			continue
		}

		article := domain.Article{
			Sha:     file.SHA,
			Name:    strings.TrimSuffix(file.Name, ".md"),
			Path:    strings.TrimSuffix(file.Path, ".md"),
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

	return articles, nil
}
