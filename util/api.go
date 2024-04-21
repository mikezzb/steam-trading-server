package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context, pageSize int) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil && page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
