package main

import (
	"fmt"
	"log"
	"os"

	"github.com/almarufh/backend-shortlink/internal/config"
	"github.com/almarufh/backend-shortlink/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	// ENV
	conf := config.LoadEnv()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST_SWAGG"), os.Getenv("PORT_SWAGG"))

	// Redis
	redis := conf.REDIS.Connect()
	defer func() {
		if err := redis.Close(); err != nil {
			log.Println("Close Redis : ", err.Error())
		}
	}()

	// PostgreSQL
	pool := conf.POSTGRES.Connect()
	defer pool.Close()

	// Gin Gonic
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	router.NewRouter(server, pool, redis, conf.ALLOWED_ORIGINS).Connect()

	server.Run(fmt.Sprintf("%s:%s", conf.HOST, conf.PORT))
}
