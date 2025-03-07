package model

type CategoryModel struct {
	Model
	Name        string `json:"name,omitempty" validate:"required,max=50,min=1"`
	Description string `json:"description,omitempty" validate:"required,max=255,min=6"`
	HexColor    string `json:"hex_color,omitempty" validate:"required,max=7,min=7"`
	NerdIcon    string `json:"nerd_icon,omitempty" validate:"required,max=10,min=1"`
}

type CategoryResponse struct {
	Page  string           `json:"page,omitempty"`
	Count string           `json:"count"`
	Data  []*CategoryModel `json:"data"`
}
