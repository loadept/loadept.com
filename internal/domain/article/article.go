package article

import "time"

type Article struct {
	Name      string    `json:"name,omitempty" validate:"required,max=36,min=1"`
	Path      string    `json:"path,omitempty"`
	Sha       string    `json:"sha,omitempty"`
	HtmlURL   string    `json:"html_url,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleResponse struct {
	Articles []Article `json:"articles"`
}
