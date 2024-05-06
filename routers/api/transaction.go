package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get transaction
// @Security Bearer
// @Param name query string true "Item Name"
// @Param days query int true "Days"
// @Param page query int true "Page"
// @Param rarity query string false "Item Rarity"
// @Param paintSeed query string false "Item Paint Seed"
// @Produce json
// @Router /transactions [get]
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func GetTransactions(c *gin.Context) {
	appG := app.Gin{C: c}

	dayStr := c.Query("days")

	// filter by days
	if dayStr != "" {
		days, err := strconv.Atoi(dayStr)
		if err != nil {
			appG.Error(http.StatusBadRequest, e.INVALID_PARAMS, fmt.Errorf("days must be a number"))
			return
		}

		transactionService := services.Transaction{}

		// get query params as map
		itemFilters := util.NewItemFilters(c.Request.URL.Query())
		transactions, err := transactionService.GetTransactionsByDays(days, itemFilters)
		if err != nil {
			appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
			return
		}

		data := make(map[string]interface{})
		data["transactions"] = transactions

		appG.Response(http.StatusOK, e.SUCCESS, data)
		return
	}

	pageStr := c.Query("page")
	if pageStr == "" {
		appG.Error(http.StatusBadRequest, e.INVALID_PARAMS, fmt.Errorf("page is required"))
		return
	}

	// filter by page
	page := util.GetPage(c)
	transactionService := services.Transaction{}
	// get query params as map
	itemFilters := util.NewItemFilters(c.Request.URL.Query())
	transactions, err := transactionService.GetTransactionsByPage(page, itemFilters)
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := transactionService.Count(itemFilters)
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := util.MakePagingData(page, setting.App.TransactionPageSize, total)
	data["transactions"] = transactions

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
