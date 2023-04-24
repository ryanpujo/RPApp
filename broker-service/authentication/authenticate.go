package authentication

import (
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type authentication struct {
	authClient *auth.Client
}

func NewAuthentication(auth *auth.Client) *authentication {
	return &authentication{auth}
}

func (a *authentication) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
			return
		}
		idToken := getTokenFromAuthHeader(authHeader)

		// verify the token
		_, err := a.authClient.VerifyIDToken(c, idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
			return
		}

		// continue to the next handler
		c.Next()
	}
}

func getTokenFromAuthHeader(header string) string {
	prefix := "Bearer "
	if len(header) > len(prefix) && header[:len(prefix)] == prefix {
		return strings.Clone(header[len(prefix):])
	}
	return ""
}
