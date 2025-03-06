package repository

import (
	"database/sql"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
)

type ArticleRepository struct {
	con *sql.DB
}

func NewArticleRepository(con *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		con: con,
	}
}

func (a *ArticleRepository) Register(model *model.ArticleModel) error {
	query := `
	INSERT INTO articles
	(user_id, title, description, content, category_id, published)
	VALUES (?, ?, ?, ?, ?, ?);
	`
	stmt, err := a.con.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		model.UserID,
		model.Title,
		model.Description,
		model.Content,
		model.CategoryID,
		model.Published,
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

func (a *ArticleRepository) SelectArticle(articleID int) (*model.ArticleModel, error) {
	query := `
	SELECT
		users.username
		articles.title,
		articles.content,
		articles.created_at
	FROM articles
	JOIN users ON users.id = articles.user_id
	WHERE article.id = ?;
	`
	stmt, err := a.con.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows := stmt.QueryRow(articleID)

	article := &model.ArticleModel{}
	if err := rows.Scan(
		&article.Title,
		&article.Description,
		&article.CreatedAt,
	); err != nil {
		return nil, err
	}

	return article, nil
}

func (a *ArticleRepository) SelectAllArticles(categoryName string, limit, page int) ([]*model.ArticleModel, error) {
	query := `
	SELECT
		articles.id,
		articles.title,
		articles.description,
		articles.updated_at
	FROM articles
	JOIN users ON users.id = articles.user_id
	JOIN categories ON categories.id = category_id
	WHERE published = false AND categories.name LIKE ?
	ORDER BY articles.updated_at DESC
	LIMIT ? OFFSET ?;
	`
	offset := (page - 1) * limit

	stmt, err := a.con.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryName, limit, offset)
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
			&article.CreatedAt,
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
