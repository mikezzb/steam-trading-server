package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		return 1
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

func HashPassword(password string) (string, error) {
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwdBytes), nil
}

func ComparePassword(hashedPwd, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

func MakePagingData(page, pageSize int, total int64) map[string]interface{} {
	return map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
	}
}
