package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	er "github.com/spriigan/broker/pkg/error"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/domain"
	"google.golang.org/api/option"
)

type Authentication struct {
	AuthClient AuthClient
}

type Authenticator interface {
	Authenticate() gin.HandlerFunc
	CreateUser(c *gin.Context)
}

type IdTokenVerifier interface {
	VerifyIDToken(ctx context.Context, id string) (*auth.Token, error)
}

type AuthClient interface {
	IdTokenVerifier
	CreateUser(ctx context.Context, user *auth.UserToCreate) (*auth.UserRecord, error)
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

func (a *Authentication) CreateUser(c *gin.Context) {
	var json domain.UserToCreate
	if err := c.ShouldBindJSON(&json); err != nil {
		er.Handle(c, err)
		return
	}

	userToCreate := auth.UserToCreate{}
	userToCreate.DisplayName(fmt.Sprintf("%s %s", json.FirstName, json.LastName))
	userToCreate.Email(json.Email)
	userToCreate.EmailVerified(false)
	userToCreate.Password(json.Password)
	userToCreate.Disabled(false)
	ctx, cancel := context.WithTimeout(c, time.Second*2)
	defer cancel()
	userRecord, err := a.AuthClient.CreateUser(ctx, &userToCreate)
	if err != nil {
		er.Handle(c, err)
	}
	var res response.JsonRes
	res.CreatedUser = *userRecord
	c.JSON(http.StatusCreated, res)
}

func (a *Authentication) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var res response.JsonRes
		res.Error = "unauthorized"
		// get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		idToken := getTokenFromAuthHeader(authHeader)

		// verify the token
		_, err := a.AuthClient.VerifyIDToken(c, idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
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
