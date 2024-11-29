package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/auth/utils"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/database"
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Sessions = map[string]models.Session{}

func AuthenticateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := utils.GetTokenFromRequest(c)
		token, err := jwt.ParseWithClaims(tokenStr, &models.Payload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, _ := token.Claims.(*models.Payload)

		userData := Sessions[claims.Session]
		if userData.ExpiryTime.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "expired session"})
			return
		}

		var user models.User
		tx := database.DBConn.Where("id=?", userData.Uid).Find(&user)
		if tx.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": tx.Error.Error()})
			return
		}

		c.Set("authorizedUser", user)

		c.Next()
	}
}
