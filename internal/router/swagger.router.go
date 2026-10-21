package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) Swagger() {
	r.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
