package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
)

// @Summary Create user
// @Accept json
// @Produce json
// @Param userForm body services.SignupForm true "User Form"
// @Router /users [post]
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	form := &services.SignupForm{}

	httpCode, errCode := app.BindValidate(c, form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := services.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
	}

	token, err := userService.CreateUser()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["token"] = token

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
