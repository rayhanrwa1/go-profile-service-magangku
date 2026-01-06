package middleware

import (
	"net/http"

	"go-profile-service-magangku/internal/response"

	"github.com/gin-gonic/gin"
)

func UserOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("userable_type")

		if role != "user" {
			c.AbortWithStatusJSON(http.StatusForbidden, response.APIResponse{
				Message: "only user can access this resource",
				Data:    nil,
			})
			return
		}

		c.Next()
	}
}
