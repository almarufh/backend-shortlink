package router

import (
	"github.com/almarufh/backend-shortlink/internal/controller"
	"github.com/almarufh/backend-shortlink/internal/repository"
	"github.com/almarufh/backend-shortlink/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Router struct {
	router         *gin.Engine
	pool           *pgxpool.Pool
	redis          *redis.Client
	allowedOrigins []string
	as             *controller.AuthContorller
}

func NewRouter(router *gin.Engine, pool *pgxpool.Pool, redis *redis.Client, allowedOrigins []string) *Router {

	ar := repository.NewAuthRepo(pool)
	as := services.NewAuthService(ar, redis)
	ac := controller.NewAuthContorller(as)

	return &Router{
		router:         router,
		pool:           pool,
		redis:          redis,
		allowedOrigins: allowedOrigins,
		as:             ac,
	}
}

func (r *Router) Connect() {
	r.StatusServer()
	r.Swagger()
	r.AuthRouter()
}
