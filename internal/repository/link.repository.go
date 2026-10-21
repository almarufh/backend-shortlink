package repository

import (
	"context"
	"errors"
	"log"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepo struct {
	db *pgxpool.Pool
}

func NewLinkRepo(db *pgxpool.Pool) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

func (lr *LinkRepo) CreateShortURLrepo(c context.Context, body dto.LinksRequest) (err error) {
	query := `
		INSERT INTO links (user_id, original_url, short_url)
		VALUES ($1, $2, $3)
	`
	args := []any{body.UserID, body.OriginalURL, body.ShortURL}

	tag, err := lr.db.Exec(c, query, args...)
	if err != nil {
		log.Printf("[LinkRepo.RegisterAR] Exec : %s", err)
		var repoErr *pgconn.PgError
		if errors.As(err, &repoErr) {
			code, msg := repoErr.Code, repoErr.ConstraintName
			log.Printf("code : %s\nmessage : %s", code, msg)
			if code == "23505" && msg == "links_short_url_key" {
				log.Printf("code : %s\nmessage : %s cek", code, msg)
				return message.ErrShortLinkExist
			}
		}
		return err
	}

	if tag.RowsAffected() == 0 {
		return message.ErrNoRowAffected
	}
	return nil
}

func (lr *LinkRepo) GetOriginalURLrepo(c context.Context, shortURL string) (*string, error) {
	query := `
		SELECT original_url
		FROM links
		WHERE short_url = $1
	`
	var links string
	err := lr.db.QueryRow(c, query, shortURL).Scan(
		&links,
	)

	if err != nil {
		log.Printf("[AuthRepo.GetUserByEmail] QueryRow : %s", err)
		return nil, message.ErrShortLinkNotExist
	}

	return &links, nil
}
