package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebastianMurdoch/go-users-example/domain"
)

type UsersRepositoryImpl struct {
	DB *sqlx.DB `inject:"auto"`
}

func (r *UsersRepositoryImpl) FindAll() []domain.User {
	allUsers := []domain.User{}
	r.DB.Select(&allUsers, "SELECT * FROM users")
	return allUsers
}

func (r *UsersRepositoryImpl) Save(user domain.User) error{
	var id int
	getNextID := `SELECT id FROM users ORDER BY id DESC LIMIT 1`
	row := r.DB.QueryRow(getNextID)
	row.Scan(&id)
	id += 1
	insertQuery := `INSERT INTO users (id, username) VALUES ($1, $2)`
	_, err := r.DB.Exec(insertQuery, id, user.Username)
	return err
}
