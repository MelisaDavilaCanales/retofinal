package persons

import "github.com/gin-gonic/gin"

func AddPersonsRoutes(r *gin.Engine) {
	persons := r.Group("/persons")

	persons.GET("/", GetPersons)
}
