package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get listings, with JWT bearer token auth
// @Produce json
// @Param page query int false "Page"
// @Router /listings [get]
func GetListings(c *gin.Context) {
	appG := app.Gin{C: c}

	listingService := services.Listing{
		Page: util.GetPage(c),
	}

	listings, err := listingService.GetListings()

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := listingService.Count()

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := make(map[string]interface{})
	data["listings"] = listings
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
