package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/types"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get transaction
// @Security Bearer
// @Param name query string true "Item Name"
// @Param page query int true "Page"
// @Param rarity query string false "Item Rarity"
// @Param paintSeed query string false "Item Paint Seed"
// @Produce json
// @Router /transactions [get]
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func GetTransactions(c *gin.Context) {
	appG := app.Gin{C: c}

	// get query params as map
	itemFilters := types.NewItemFilters(c.Request.URL.Query())

	transactionService := services.Transaction{
		ItemFilters: itemFilters,
		Page:        util.GetPage(c),
	}

	transactions, err := transactionService.GetTransactions()
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := transactionService.Count()
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := make(map[string]interface{})
	data["transactions"] = transactions
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
