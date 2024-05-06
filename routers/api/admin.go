package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get items [admin]
// @Security Bearer
// @Produce json
// @Param page query int false "Page"
// @Router /admin/items [get]
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func GetItemsAdmin(c *gin.Context) {
	appG := app.Gin{C: c}

	itemService := services.Item{
		PageNum: util.GetPage(c),
	}

	items, err := itemService.GetItems(100, nil)
	itemFilters := util.NewItemFilters(c.Request.URL.Query())

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := itemService.Count(itemFilters)
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := make(map[string]interface{})
	data["items"] = items
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
