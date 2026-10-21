package middleware

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	allowedOrigins []string
	pool           *pgxpool.Pool
	redis          *redis.Client
}

func InitMiddleware(pool *pgxpool.Pool, redis *redis.Client, allowedOrigins []string) *Middleware {
	return &Middleware{
		allowedOrigins: allowedOrigins,
		pool:           pool,
		redis:          redis,
	}
}
