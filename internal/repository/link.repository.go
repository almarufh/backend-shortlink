package repository

import (
	"context"
	"errors"
	"log"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/almarufh/backend-shortlink/internal/models"
	"github.com/jackc/pgx/v5"
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
		WHERE short_url = $1 AND deleted_at IS NULL
	`
	var links string
	err := lr.db.QueryRow(c, query, shortURL).Scan(
		&links,
	)

	if err != nil {
		log.Printf("[GetOriginalURLrepo] QueryRow : %s", err)
		return nil, message.ErrShortLinkNotExist
	}

	return &links, nil
}

func (lr *LinkRepo) GetMyLinksRepo(c context.Context, userId int32) (*[]models.Links, error) {
	query := `
		SELECT 
		    id,
		    original_url,
		    short_url
		FROM links
		WHERE user_id = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := lr.db.Query(c, query, userId)
	if err != nil {
		log.Printf("[GetMyLinksRepo] Query : %s", err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, message.ErrInternal
	}

	defer rows.Close()

	links := make([]models.Links, 0)

	for rows.Next() {
		var link models.Links
		err := rows.Scan(
			&link.ID,
			&link.OriginalURL,
			&link.ShortURL,
		)
		if err != nil {
			log.Printf("[GetMyLinksRepo] rows.Next : %s", err.Error())
			return nil, message.ErrFailedQuery
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[GetMyLinksRepo] rows.Err : %s", err.Error())
		return nil, message.ErrReadRow
	}
	return &links, nil
}
