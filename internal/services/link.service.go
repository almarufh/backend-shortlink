package services

import (
	"context"
	"errors"

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
