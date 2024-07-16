package repository

import (
	"database/sql"
	"github.com/yuvakkrishnan/user-service/internal/models"
)

type UserRepository interface {
    Create(user *models.User) error
    GetByUsername(username string) (*models.User, error)
    GetByID(userID int64) (*models.User, error)
}

type userRepo struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) Create(user *models.User) error {
    query := `INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
    _, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
    return err
}

func (r *userRepo) GetByUsername(username string) (*models.User, error) {
    query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE username = $1`
    row := r.db.QueryRow(query, username)
    user := &models.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *userRepo) GetByID(userID int64) (*models.User, error) {
    query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = $1`
    row := r.db.QueryRow(query, userID)
    user := &models.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return user, nil
}
