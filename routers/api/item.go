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

// @Summary Get items
// @Produce json
// @Param page query int false "Page"
// @Router /items [get]
func GetItems(c *gin.Context) {
	appG := app.Gin{C: c}

	itemService := services.Item{
		PageNum: util.GetPage(c, setting.App.ItemPageSize),
	}

	items, err := itemService.GetItems()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	// add total count
	total, err := itemService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = items
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
