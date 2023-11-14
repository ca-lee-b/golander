package repositories

import (
	"database/sql"

	"github.com/ca-lee-b/golander/server/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetOneById(id string) *models.User {
	var user models.User
	row := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepository) CreateUser(email string, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO users(email, password) VALUES ($1, $2)", email, hashed)
	if err != nil {
		return err
	}

	return nil
}
