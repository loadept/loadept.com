package model

type CategoryModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty" validate:"required,max=50,min=1"`
	HexColor string `json:"hex_color,omitempty" validate:"required,hexcolor"`
	NerdIcon string `json:"nerd_icon,omitempty" validate:"required,max=10,min=1"`
}

type CategoryResponse struct {
	Category []CategoryModel `json:"categories"`
}
