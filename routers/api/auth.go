package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/app"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/services"
	"github.com/mikezzb/steam-trading-server/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Post auth
// @Accept json
// @Produce json
// @Param loginForm body services.LoginForm true "Login Form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func PostAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	form := &services.LoginForm{}
	httpCode, errCode := app.BindValidate(c, form)

	log.Printf("form: %v", form)
	log.Printf("post body: %v", c.Request.Body)

	if errCode != e.SUCCESS {
		appG.Error(httpCode, errCode, nil)
		return
	}

	// check auth
	user := &services.User{Email: form.Email, Password: form.Password}
	isExist, err := user.CheckAuth()
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.ERROR_USER_WRONG_PWD, err)
		return
	}

	if !isExist {
		appG.Error(http.StatusUnauthorized, e.ERROR_USER_NOT_EXIST, err)
		return
	}

	token, err := util.GenerateToken(user.User.ID.Hex(), user.User.Role)
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.SERVER_ERROR, err)
		return
	}

	data := make(map[string]interface{})
	data["token"] = token
	data["user"] = user.GetViewableUser()

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary Get auth
// @Security Bearer
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}

	// find user by id
	userId, ok := c.Get("userId")
	if !ok {
		appG.Error(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	user := &services.User{ID: userId.(primitive.ObjectID)}

	err := user.LoadUser()
	if err != nil {
		appG.Error(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, err)
		return
	}

	data := make(map[string]interface{})
	data["user"] = user.GetViewableUser()

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
