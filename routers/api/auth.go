package api

import (
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username" form:"username" query:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password" form:"password" query:"password"`
}

// @Summary Get auth
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func PostAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	vali := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	a := auth{Username: username, Password: password}
	ok, _ := vali.Valid(&a)
	if !ok {
		app.MakeErrors(vali.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// check auth
	user := &services.User{Username: username, Password: password}
	isExist, err := user.CheckAuth()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_USER_NOT_EXIST, nil)
		return
	}

	token, err := util.GenerateToken(user.Username, user.User.ID.Hex())
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
