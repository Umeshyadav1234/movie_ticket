package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"MTBS/jwt"

	"github.com/gin-gonic/gin"
)

func AuthorizeRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if role is in allowedRoles
		authorized := false
		for _, allowed := range allowedRoles {
			for _, actual := range claims.Role {
				if string(actual) == allowed {
					authorized = true
					break
				}
			}
			if authorized {
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("privileges", claims.Privileges)
		c.Next()
	}
}
func RequirePrivilege(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRole, _ := c.Get("role")
		tokenPrivs, exists := c.Get("privileges")

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Privileges missing"})
			c.Abort()
			return
		}

		privileges := tokenPrivs.([]string)
		for _, p := range privileges {
			if p == required {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": fmt.Sprintf("'%s' role does not have '%s' privilege", tokenRole, required),
		})
		c.Abort()
	}
}
