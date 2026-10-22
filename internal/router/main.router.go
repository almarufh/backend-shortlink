package router

import (
	"github.com/almarufh/backend-shortlink/internal/controller"
	"github.com/almarufh/backend-shortlink/internal/middleware"
	"github.com/almarufh/backend-shortlink/internal/repository"
	"github.com/almarufh/backend-shortlink/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Router struct {
	home           *gin.Engine
	router         *gin.RouterGroup
	pool           *pgxpool.Pool
	redis          *redis.Client
	allowedOrigins []string
	m              *middleware.Middleware
	ac             *controller.AuthContorller
	lc             *controller.LinksControl
}

func NewRouter(router *gin.Engine, pool *pgxpool.Pool, redis *redis.Client, allowedOrigins []string) *Router {
	m := middleware.InitMiddleware(pool, redis, allowedOrigins)
	router.Use(m.Cors)
	api := router.Group("api")

	ar := repository.NewAuthRepo(pool)
	as := services.NewAuthService(ar, redis)
	ac := controller.NewAuthContorller(as)

	lr := repository.NewLinkRepo(pool)
	ls := services.NewLinksService(lr)
	lc := controller.NewLinksControl(ls)

	return &Router{
		router:         api,
		pool:           pool,
		redis:          redis,
		allowedOrigins: allowedOrigins,
		m:              m,
		ac:             ac,
		lc:             lc,
		home:           router,
	}
}

func (r *Router) Connect() {
	r.Swagger()
	r.AuthRouter()
	r.LinksRouter()
}
