package HTTP_API

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

/*
Middleware for checking api key/token for each request.
*/
func (a *Controller) checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtString string
		var err error

		tokenHeader := c.GetHeader("Authorization")
		if len(tokenHeader) > 0 {
			headerArr := strings.Split(tokenHeader, " ")
			if len(headerArr) != 2 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			jwtString = headerArr[1]
		} else {
			jwtString, err = c.Cookie("JWT")
			if err != nil || jwtString == "" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatus(http.StatusUnauthorized)
				return nil, errors.New("api token is not valid")
			}
			return a.jwtSecret, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId, err := primitive.ObjectIDFromHex(claims["id"].(string))
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("userId", userId)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
