// middleware/auth.go

package middleware

import (
	"gumuruh-clinic/config"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserFromToken extracts user information from JWT token
func GetUserFromToken(c *gin.Context) (*jwt.Token, jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, nil, nil
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Use the secret from config
		return []byte(config.LoadConfig().JWTSecret), nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, err
	}

	return token, claims, nil
}

// RequireAuth verifies JWT token
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, claims, err := GetUserFromToken(c)
		if err != nil || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", claims["user_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}

// RequireAdmin ensures user is an admin
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireAdminOrSelf allows admin or the user themselves
func RequireAdminOrSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		userID, _ := c.Get("userID")
		paramID := c.Param("id")

		if role == "admin" {
			c.Next()
			return
		}

		if userID.(string) == paramID {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
	}
}

// RequireAdminOrDoctor allows admin or doctor access
func RequireAdminOrDoctor() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role != "admin" && role != "dokter") {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin or Doctor access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
