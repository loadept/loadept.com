package repository

import (
	"database/sql"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
)

type CategoryRepository struct {
	conn *sql.DB
}

func NewCetogoryRepository(conn *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		conn: conn,
	}
}

func (a *CategoryRepository) RegisterCategory(model *model.CategoryModel) error {
	query := `
	INSERT INTO categories 
	(id, name, description, hex_color, utf_icon)
	VALUES (?, ?, ?, ?, ?);`
	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		model.ID,
		model.Name,
		model.Description,
		model.HexColor,
		model.NerdIcon,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("Expected 1 row affected, got %d", rowsAffected)
	}

	return nil
}

func (a *CategoryRepository) GetCategoryByID(categoryID string) (*model.CategoryModel, error) {
	query := `
	SELECT
		name,
		description,
		hex_color,
		utf_icon
	FROM categories
	WHERE id = ?;`
	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(categoryID)

	category := &model.CategoryModel{}
	if err := rows.Scan(
		&category.Name,
		&category.Description,
		&category.HexColor,
		&category.NerdIcon,
	); err != nil {
		return nil, err
	}

	return category, nil
}

func (a *CategoryRepository) SelectCategories(limit, page int) ([]*model.CategoryModel, error) {
	query := `
	SELECT
		id,
		name,
		description,
		hex_color,
		utf_icon
	FROM categories
	LIMIT ? OFFSET ?;`

	offset := (page - 1) * limit

	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*model.CategoryModel

	for rows.Next() {
		category := &model.CategoryModel{}
		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.HexColor,
			&category.NerdIcon,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if categories == nil {
		categories = make([]*model.CategoryModel, 0)
	}
	return categories, nil
}
