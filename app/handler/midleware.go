package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/hendrihmwn/dating-app-api/app/config"
	"net/http"
	"strconv"
	"strings"
)

func TokenAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")

	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "invalid authorizaion header",
		})
		return
	}

	tokenString := strings.ReplaceAll(authorizationHeader, "Bearer ", "")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(config.GetAccessTokenSecretKey()), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "invalid authorizaion header",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "invalid authorizaion header",
		})
		return
	}

	userId, err := strconv.Atoi(fmt.Sprint(claims["id"]))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "invalid token",
		})
		return
	}

	c.Set("user_id", userId)
	c.Next()
}
