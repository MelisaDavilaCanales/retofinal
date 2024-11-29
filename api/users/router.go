package users

import (
	"github.com/MelisaDavilaCanales/trucode3-challenge-final.git/auth"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.Use(auth.AuthenticateSession())
	group.GET("/me", GetMe)
}
