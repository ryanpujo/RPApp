package authentication

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type Authentication struct {
	AuthClient AuthClient
}

type Authenticator interface {
	Authenticate() gin.HandlerFunc
}

type AuthClient interface {
	VerifyIDToken(ctx context.Context, id string) (*auth.Token, error)
}

func NewAuthentication() *Authentication {
	config := firebase.Config{
		ProjectID: "orbit-app-145b9",
	}
	opt := option.WithCredentialsFile("./orbit-app-145b9-firebase-adminsdk-7ycvp-6ab97f8272.json")
	app, err := firebase.NewApp(context.Background(), &config, opt)
	if err != nil {
		log.Fatal(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return &Authentication{AuthClient: authClient}
}

func (a *Authentication) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unathorized"})
			return
		}
		idToken := getTokenFromAuthHeader(authHeader)

		// verify the token
		_, err := a.AuthClient.VerifyIDToken(c, idToken)
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
