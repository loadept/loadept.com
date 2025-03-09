package repository

import (
	"database/sql"
	"fmt"

	"github.com/loadept/loadept.com/internal/model"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) RegisterUser(model *model.RegisterUserModel) error {
	query := `
	INSERT INTO users
	(id, full_name, username, email, password)
	VALUES (?, ?, ?, ?, ?);`
	stmt, err := u.conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		model.ID,
		model.FullName,
		model.Username,
		model.Email,
		model.Password,
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

func (u *UserRepository) GetUserByName(username string) (*model.UserModel, error) {
	query := `
	SELECT id, username, is_admin, password
	FROM users WHERE username = ?;`
	stmt, err := u.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	user := &model.UserModel{}
	if err = row.Scan(
		&user.ID,
		&user.Username,
		&user.IsAdmin,
		&user.Password,
	); err != nil {
		return nil, err
	}

	return user, nil
}
