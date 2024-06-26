package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/db"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/routers"
	"github.com/mikezzb/steam-trading-server/services"
)

func init() {
	setting.Setup()
	db.Setup()
	// must be after db.Setup() to use Repos
	services.Setup()
}

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	gin.SetMode(setting.Server.RunMode)

	router := routers.InitRouter()

	var endPoint string
	if os.Getenv("PORT") != "" {
		endPoint = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		endPoint = fmt.Sprintf(":%d", setting.Server.HttpPort)
	}

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    setting.Server.ReadTimeout,
		WriteTimeout:   setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] http server listening %s", endPoint)

	server.ListenAndServe()
}
