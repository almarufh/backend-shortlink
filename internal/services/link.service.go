package services

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/almarufh/backend-shortlink/internal/repository"
	"github.com/almarufh/backend-shortlink/internal/utils"
)

type LinksService struct {
	lr *repository.LinkRepo
}

func NewLinksService(lr *repository.LinkRepo) *LinksService {
	return &LinksService{
		lr: lr,
	}
}

func (ls *LinksService) CreateShortURLservice(c context.Context, body dto.LinksRequest) error {
	if body.ShortURL != nil && *body.ShortURL != "" {
		return ls.lr.CreateShortURLrepo(c, body)
	}

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		slug, err := utils.GenerateRandomSlug(6)
		if err != nil {
			return err
		}

		body.ShortURL = &slug

		err = ls.lr.CreateShortURLrepo(c, body)
		if err == nil {
			return nil
		}

		if !errors.Is(err, message.ErrShortLinkExist) {
			return err
		}
	}

	return message.ErrShortLinkExist
}

func (ls *LinksService) GetOriginalURLService(ctx context.Context, shortURL string) (*string, error) {
	link, err := ls.lr.GetOriginalURLrepo(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (ls *LinksService) GetMyLinksService(ctx context.Context, userID int32) (*dto.MyLinksResponse, error) {
	links, err := ls.lr.GetMyLinksRepo(ctx, userID)
	if err != nil {
		return nil, err
	}

	results := make([]dto.Links, 0, len(*links))

	for _, link := range *links {
		slug := fmt.Sprintf("http://%s:%s/%s", os.Getenv("HOST_SWAGG"), os.Getenv("PORT_SWAGG"), link.ShortURL)
		results = append(results, dto.Links{
			LinksID:     link.ID,
			OriginalURL: link.OriginalURL,
			ShortURL:    slug,
			Slug:        link.ShortURL,
		})
	}

	return &dto.MyLinksResponse{
		Data: results,
	}, nil
}
