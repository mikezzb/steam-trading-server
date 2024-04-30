package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mikezzb/steam-trading-server/e"
	"github.com/mikezzb/steam-trading-server/util"
)

func JWT() gin.HandlerFunc {
	return JWTRole("")
}

func JWTRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims = nil

		code = e.SUCCESS
		// extract token from auth header

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			code = e.ERROR_INVALID_AUTH_HEADER
		}

		authHeaderParts := strings.Split(authHeader, " ") // Bearer token
		log.Printf("authHeaderParts: %v", authHeaderParts)
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			code = e.ERROR_INVALID_AUTH_HEADER
		} else {

			token := authHeaderParts[1]

			parsedClaims, err := util.ParseToken(token)
			claims = parsedClaims

			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_EXPIRED
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}

			// check role
			if role != "" && claims.Role != role {
				code = e.ERROR_AUTH_CHECK_ROLE_FAIL
			}
		}

		// if error, return
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		// add claims to context
		c.Set("role", claims.Role)
		c.Set("userId", util.StringToObjectId(claims.UserId))

		c.Next()
	}
}
