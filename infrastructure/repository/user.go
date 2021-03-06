package repository

import (
	"context"
	"database/sql"
	"os"

	"github.com/lib/pq"
	"github.com/pocket7878/spa_login_learning_backend/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepositoryImpl, error) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		return nil, err
	}

	return &UserRepositoryImpl{db}, nil
}

func (u *UserRepositoryImpl) GetByProviderWithUID(ctx context.Context, provider, uid string) (*domain.User, error) {
	query := `SELECT id,provider,uid FROM users WHERE provider=$1 AND uid=$2`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, provider, uid)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	if rows.Next() == true {
		result := new(domain.User)
		err = rows.Scan(&result.ID, &result.Provider, &result.UID)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return nil, nil
	}
}

func (u *UserRepositoryImpl) Store(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (provider, uid) VALUES ($1, $2) RETURNING id`
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, user.Provider, user.UID).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
