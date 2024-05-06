package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get listings
// @Security Bearer
// @Produce json
// @Param page query int false "Page"
// @Param name query string true "Item Name"
// @Param page query int true "Page"
// @Param rarity query string false "Item Rarity"
// @Param paintSeed query string false "Item Paint Seed"
// @Router /listings [get]
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func GetListings(c *gin.Context) {
	appG := app.Gin{C: c}

	listingService := services.Listing{}

	itemFilters := util.NewItemFilters(c.Request.URL.Query())
	page := util.GetPage(c)

	listings, err := listingService.GetListingsByPage(page, itemFilters)

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := listingService.Count(itemFilters)

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := util.MakePagingData(page, setting.App.ListingPageSize, total)
	data["listings"] = listings

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
