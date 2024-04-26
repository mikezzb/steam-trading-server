package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPage(c *gin.Context, pageSize int) int {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		return 0
	}
	return page
}

func ObjectIdToString(id primitive.ObjectID) string {
	if id.IsZero() {
		return ""
	}
	return id.Hex()
}

func StringToObjectId(id string) primitive.ObjectID {
	if id == "" {
		return primitive.NilObjectID
	}
	objId, _ := primitive.ObjectIDFromHex(id)
	return objId
}

func GetUserId(c *gin.Context) primitive.ObjectID {
	userId, _ := c.Get("userId")
	return userId.(primitive.ObjectID)
}
