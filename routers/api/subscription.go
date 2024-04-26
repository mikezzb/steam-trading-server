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
// @Produce json
// @Router /subscription [post]
func AddSub(c *gin.Context) {
	appG := app.Gin{C: c}
	var form *services.AddSubForm = nil

	httpCode, errCode := app.BindValidate(c, form)
	if errCode != e.SUCCESS || form == nil {
		appG.Response(httpCode, errCode, nil)
		return
	}

	subService := services.Subscription{
		Name:       form.Name,
		Rarity:     form.Rarity,
		MaxPremium: form.MaxPremium,
		NotiType:   form.NotiType,
		NotiId:     form.NotiId,
		OwnerId:    util.GetUserId(c),
	}

	id, err := subService.AddSub()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["id"] = id

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func UpdateSub(c *gin.Context) {
	appG := app.Gin{C: c}
	var form *services.UpdateSubForm = nil

	httpCode, errCode := app.BindValidate(c, form)
	if errCode != e.SUCCESS || form == nil {
		appG.Response(httpCode, errCode, nil)
		return
	}

	subService := services.Subscription{
		ID:         form.ID,
		Name:       form.Name,
		Rarity:     form.Rarity,
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

func DeleteSub(c *gin.Context) {
	appG := app.Gin{C: c}
	id := c.Query("id")

	if id == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	subService := services.Subscription{
		ID:      id,
		OwnerId: util.GetUserId(c),
	}

	err := subService.DeleteSub()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetSubs(c *gin.Context) {
	appG := app.Gin{C: c}

	subService := services.Subscription{
		OwnerId: util.GetUserId(c),
	}

	subs, err := subService.GetSubs()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["subscriptions"] = subs

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
