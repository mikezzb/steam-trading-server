package api

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Get items
// @Produce json
// @Param page query int true "Page"
// @Router /items [get]
func GetItems(c *gin.Context) {
	appG := app.Gin{C: c}

	itemService := services.Item{
		PageNum: util.GetPage(c),
	}

	log.Printf("GetItems: %d", itemService.PageNum)

	items, err := itemService.GetItems(setting.App.ItemPageSize)

	runtime.Gosched()

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	// add total count
	total, err := itemService.Count()
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := util.MakePagingData(itemService.PageNum, setting.App.ItemPageSize, total)
	data["items"] = items

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary Get item by id
// @Produce json
// @Param id path string true "Item ID"
// @Router /items/{id} [get]
func GetItem(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Param("id")

	itemService := services.Item{}
	item, err := itemService.GetItem(id)

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["item"] = item
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
