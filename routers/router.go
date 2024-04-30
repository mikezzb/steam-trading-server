package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/mikezzb/steam-trading-server/docs"
	"github.com/mikezzb/steam-trading-server/middleware"
	"github.com/mikezzb/steam-trading-server/pkg/static"
	"github.com/mikezzb/steam-trading-server/routers/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	// swagger
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// static files
	r.StaticFS("/static/images", http.Dir(static.GetImageFullUrl()))

	// api
	apiv1 := r.Group("/api")
	{
		// auth
		apiv1.POST("/auth", api.PostAuth)
		apiv1.GET("/auth", middleware.JWT(), api.GetAuth)

		// items
		apiv1.GET("/items", api.GetItems)
		apiv1.GET("/items/:id", api.GetItem)

		// listings
		apiv1.GET("/listings", middleware.JWT(), api.GetListings)

		// subscriptions
		apiv1.GET("/subscriptions", middleware.JWT(), api.GetSubs)
		apiv1.POST("/subscriptions", middleware.JWT(), api.AddSub)
		apiv1.DELETE("/subscriptions/:id", middleware.JWT(), api.DeleteSub)
		apiv1.PUT("/subscriptions/:id", middleware.JWT(), api.UpdateSub)

		// transactions
		apiv1.GET("/transactions", middleware.JWT(), api.GetTransactions)

		// users
		apiv1.POST("/users", api.CreateUser)

		// admins
		apiv1.GET("/admin/items", middleware.JWTRole("admin"), api.GetItemsAdmin)
	}

	return r
}
