package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/mikezzb/steam-trading-server/docs"
	"github.com/mikezzb/steam-trading-server/pkg/static"
	"github.com/mikezzb/steam-trading-server/routers/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// static files
	r.StaticFS("/static/images", http.Dir(static.GetImageFullUrl()))

	// auth
	r.POST("/auth", api.GetAuth)

	// api
	apiv1 := r.Group("/api")
	{
		// items
		apiv1.GET("/items", api.GetItems)

		// listings
	}

	return r
}
