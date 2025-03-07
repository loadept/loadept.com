package model

import "time"

type ArticleModel struct {
	UserModel   `validate:"-"`
	UserID      string     `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty" validate:"required,max=36,min=1"`
	Description string     `json:"description,omitempty" validate:"required,max=255,min=1"`
	Content     string     `json:"content,omitempty" validate:"required,min=1"`
	CategoryID  string     `json:"category_id,omitempty" validate:"uuid"`
	Published   bool       `json:"published,omitempty" validate:"boolean"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type ArticleResponse struct {
	Page  string          `json:"page,omitempty"`
	Count string          `json:"count"`
	Data  []*ArticleModel `json:"data"`
}
