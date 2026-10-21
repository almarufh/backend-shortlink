package repository

import (
	"context"
	"errors"
	"log"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/almarufh/backend-shortlink/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar *AuthRepo) RegisterAR(c context.Context, body dto.RegisterRequest) (err error) {
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`
	args := []any{body.Email, body.Password}

	tag, err := ar.db.Exec(c, query, args...)
	if err != nil {
		log.Printf("[AuthRepo.RegisterAR] Exec : %s", err)
		var repoErr *pgconn.PgError
		if errors.As(err, &repoErr) {
			code, msg := repoErr.Code, repoErr.ConstraintName
			log.Printf("code : %s\nmessage : %s", code, msg)
			if code == "23505" && msg == "users_email_key" {
				log.Printf("code : %s\nmessage : %s cek", code, msg)
				return message.ErrEmailExist
			}
		}
		return err
	}

	if tag.RowsAffected() == 0 {
		return message.ErrNoRowAffected
	}
	return nil
}

func (ar *AuthRepo) GetUserByEmail(c context.Context, body dto.LoginRequest) (*models.Users, error) {
	query := `
		SELECT
			id,
		    password
		FROM users
		WHERE email = $1
    `
	var user models.Users
	err := ar.db.QueryRow(c, query, body.Email).Scan(
		&user.ID,
		&user.Password,
	)

	if err != nil {
		log.Printf("[AuthRepo.GetUserByEmail] QueryRow : %s", err)
		return nil, message.ErrEmailNotExist
	}

	return &user, nil
}
