package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"users": []string{"Alice", "Bob", "Charlie"},
			})
		})

		userGroup.GET("/:id", func(c *gin.Context) {
			userId := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"user": gin.H{
					"id":   userId,
					"name": "User " + userId,
				},
			})
		})
	}
}
