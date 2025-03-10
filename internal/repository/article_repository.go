package repository

import (
	"database/sql"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
)

type ArticleRepository struct {
	conn *sql.DB
}

func NewArticleRepository(conn *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		conn: conn,
	}
}

func (a *ArticleRepository) RegisterArticle(userID string, model *model.ArticleModel) error {
	query := `
	INSERT INTO articles
	(id, user_id, title, description, content, category_id)
	VALUES (?, ?, ?, ?, ?, ?);`
	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		model.ID,
		userID,
		model.Title,
		model.Description,
		model.Content,
		model.CategoryID,
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

func (a *ArticleRepository) GetArticleByID(articleID string) (*model.ArticleModel, error) {
	query := `
	SELECT
		users.username,
		articles.title,
		articles.content,
		articles.updated_at
	FROM articles
	JOIN users ON users.id = articles.user_id
	WHERE articles.id = ?;`
	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(articleID)

	article := &model.ArticleModel{}
	if err := rows.Scan(
		&article.Username,
		&article.Title,
		&article.Content,
		&article.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return article, nil
}

func (a *ArticleRepository) SelectArticles(categoryName, articleTitle string, limit, page int) ([]*model.ArticleModel, error) {
	query := `
	SELECT
		articles.id,
		articles.title,
		articles.description,
		articles.updated_at
	FROM articles
	JOIN categories ON categories.id = category_id
	WHERE published = false `

	var args []any
	if len(categoryName) != 0 {
		query += "AND categories.name LIKE ?"
		args = append(args, categoryName)
	}
	if len(articleTitle) != 0 {
		query += "AND articles.title LIKE ?"
		args = append(args, "%"+articleTitle+"%")
	}

	query += `
	ORDER BY articles.updated_at DESC
	LIMIT ? OFFSET ?;`

	offset := (page - 1) * limit

	// add limit and offset params in args slice
	args = append(args, limit, offset)

	stmt, err := a.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*model.ArticleModel

	for rows.Next() {
		article := &model.ArticleModel{}
		if err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Description,
			&article.UpdatedAt,
		); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if articles == nil {
		articles = make([]*model.ArticleModel, 0)
	}
	return articles, nil
}
