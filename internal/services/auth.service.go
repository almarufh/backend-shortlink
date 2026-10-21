package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/almarufh/backend-shortlink/internal/dto"
	message "github.com/almarufh/backend-shortlink/internal/messages"
	"github.com/almarufh/backend-shortlink/internal/repository"
	"github.com/almarufh/backend-shortlink/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	conf  *pkg.ConfigHash
	redis *redis.Client
	ar    *repository.AuthRepo
}

func NewAuthService(AuthRepo *repository.AuthRepo, redis *redis.Client) *AuthService {
	return &AuthService{
		ar:    AuthRepo,
		conf:  pkg.NewConfigHash(),
		redis: redis,
	}
}

func (as *AuthService) RegisterAS(c context.Context, body dto.RegisterRequest) (err error) {
	hashPassword, err := as.conf.HashPassword(body.Password)
	if err != nil {
		log.Printf("[AuthService.RegisterAS] HashPassword %s", err)
		return err
	}

	body.Password = hashPassword

	err = as.ar.RegisterAR(c, body)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) LoginAS(c context.Context, body dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := as.ar.GetUserByEmail(c, body)
	if err != nil {
		return nil, err
	}

	_, err = as.conf.Compare(body.Password, user.Password)
	if err != nil {
		log.Printf("[AuthService.LoginAS] Compare %s", err)
		return nil, message.ErrEmailNotExist
	}

	jwt := pkg.NewJWTClaims(user.ID, 60)

	token, err := jwt.GeneretToken()
	if err != nil {
		log.Printf("[AuthService.LoginAS] GeneretToke %s", err)
		return nil, err
	}

	key := fmt.Sprintf("shortlink:token:%d", user.ID)
	err = as.redis.Set(c, key, user.ID, 60*time.Minute).Err()
	if err != nil {
		log.Printf("[AuthService.LoginAS] Redis Set %s", err)
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
