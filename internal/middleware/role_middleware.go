package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleAuthorization(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		for _, allowed := range allowedRoles {
			if strings.EqualFold(role, allowed) {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "You don't have permission to access this resource",
		})
		c.Abort()
	}
}
