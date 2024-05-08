package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

// @Summary Add subscription
// @Security ApiKeyAuth
// @Accept json
// @Param addSubForm body services.AddSubForm true "Add Subscription Form"
// @Produce json
// @Router /subscriptions [post]
func AddSub(c *gin.Context) {
	appG := app.Gin{C: c}
	var form *services.AddSubForm = &services.AddSubForm{}

	httpCode, errCode := app.BindValidate(c, form)
	if errCode != e.SUCCESS {
		appG.Error(httpCode, errCode, nil)
		return
	}

	subService := services.Subscription{
		Name:       form.Name,
		Rarities:   form.Rarities,
		PaintSeeds: form.PaintSeeds,
		MaxPremium: form.MaxPremium,
		NotiType:   form.NotiType,
		NotiId:     form.NotiId,
		OwnerId:    util.GetUserId(c),
	}

	id, err := subService.AddSub()

	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := make(map[string]interface{})
	data["id"] = id

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Update subscription
// @Security ApiKeyAuth
// @Produce json
// @Router /subscriptions [put]
// @Param id path string true "Subscription ID"
func UpdateSub(c *gin.Context) {
	appG := app.Gin{C: c}
	var form *services.UpdateSubForm = &services.UpdateSubForm{}

	id := c.Params.ByName("id")

	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	httpCode, errCode := app.BindValidate(c, form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	subService := services.Subscription{
		ID:         id,
		Name:       form.Name,
		Rarities:   form.Rarities,
		PaintSeeds: form.PaintSeeds,
		MaxPremium: form.MaxPremium,
		NotiType:   form.NotiType,
		NotiId:     form.NotiId,

		OwnerId: util.GetUserId(c),
	}

	err := subService.UpdateSub()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Delete subscription
// @Security ApiKeyAuth
// @Produce json
// @Router /subscriptions [delete]
// @Param id path string true "Subscription ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
func DeleteSub(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Params.ByName("id")

	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	subService := services.Subscription{
		ID:      id,
		OwnerId: util.GetUserId(c),
	}

	err := subService.DeleteSub(id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Get subscriptions
// @Security ApiKeyAuth
// @Produce json
// @Router /subscriptions [get]
// @Param name query string true "Item Name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func GetSubs(c *gin.Context) {
	appG := app.Gin{C: c}

	subService := services.Subscription{
		OwnerId: util.GetUserId(c),
	}

	itemFilters := util.NewItemFilters(c.Request.URL.Query())

	subs, err := subService.GetSubs(itemFilters)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["subscriptions"] = subs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
