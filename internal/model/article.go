package model

import "time"

type ArticleModel struct {
	UserModel
	UserID      string     `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Content     string     `json:"content,omitempty"`
	CategoryID  string     `json:"category_id,omitempty"`
	Published   bool       `json:"published,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ArticleResponse struct {
	Page  string          `json:"page,omitempty"`
	Count string          `json:"count"`
	Data  []*ArticleModel `json:"data"`
}
