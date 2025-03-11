package model

import "time"

type ArticleModel struct {
	Name      string    `json:"name,omitempty" validate:"required,max=36,min=1"`
	Path      string    `json:"path,omitempty"`
	Sha       string    `json:"sha,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleResponse struct {
	Articles []ArticleModel `json:"articles"`
}
