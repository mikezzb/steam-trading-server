package app

import (
	"log"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/mikezzb/steam-trading-server/e"
)

func BindValidate(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)

	if err != nil {
		log.Printf("bind error: %v", err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		log.Printf("valid error: %v", err)
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MakeErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
