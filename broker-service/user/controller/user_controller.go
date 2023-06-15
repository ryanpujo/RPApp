package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"firebase.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	er "github.com/spriigan/broker/pkg/error"
	"github.com/spriigan/broker/response"
	"github.com/spriigan/broker/user/domain"
	"github.com/spriigan/broker/user/grpc/client"
	"github.com/spriigan/broker/user/user-proto/userpb"
)

type userController struct {
	c       client.UserClientCloser
	storage *storage.Client
}

func NewUserController(c client.UserClientCloser, s *storage.Client) *userController {
	return &userController{c: c, storage: s}
}

func (u userController) Create(c *gin.Context) {
	var json domain.User
	if err := c.ShouldBindJSON(&json); err != nil {
		er.Handle(c, err)
		return
	}

	payload := &userpb.UserPayload{
		Firstname: json.FirstName,
		Lastname:  json.LastName,
		Username:  json.Username,
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	user, err := u.c.CreateUser(ctx, payload)
	if err != nil {
		er.Handle(c, err)
		return
	}

	created := domain.User{
		ID:        user.Id,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Username:  user.Username,
		CreatedAt: sql.NullTime{Time: user.CreatedAt.AsTime()},
	}

	var res response.JsonRes
	res.User = created
	c.JSON(http.StatusCreated, res)
}

type Uri struct {
	Id int64 `uri:"id" binding:"gt=0"`
}

func (u userController) GetById(c *gin.Context) {
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	found, err := u.c.GetById(ctx, &userpb.UserId{Id: uri.Id})
	if err != nil {
		er.Handle(c, err)
		return
	}

	user := domain.User{
		ID:        found.Id,
		FirstName: found.Firstname,
		LastName:  found.Lastname,
		Username:  found.Username,
		CreatedAt: sql.NullTime{Time: found.CreatedAt.AsTime()},
	}

	var res response.JsonRes
	res.User = user
	c.JSON(http.StatusOK, res)
}

func (u userController) DeleteById(c *gin.Context) {
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	_, err := u.c.DeleteById(ctx, &userpb.UserId{Id: uri.Id})
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (u userController) UpdateById(c *gin.Context) {
	var json domain.User
	if err := c.ShouldBindJSON(&json); err != nil {
		er.Handle(c, err)
		return
	}

	payload := &userpb.UserPayload{
		Id:        json.ID,
		Firstname: json.FirstName,
		Lastname:  json.LastName,
		Username:  json.Username,
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	_, err := u.c.UpdateById(ctx, payload)
	if err != nil {
		er.Handle(c, err)
		return
	}

	c.Status(http.StatusOK)
}

type QueryParams struct {
	Limit int32 `form:"limit" binding:"gt=0"`
	Page  int32 `form:"page" binding:"gt=0"`
}

func (u userController) GetMany(c *gin.Context) {
	var Query QueryParams
	if err := c.ShouldBindQuery(&Query); err != nil {
		er.Handle(c, err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	result, err := u.c.GetMany(ctx, &userpb.GetMAnyArgs{Limit: Query.Limit, Page: Query.Page})
	if err != nil {
		er.Handle(c, err)
		return
	}

	users := make([]domain.User, 0, len(result.Users))

	for _, v := range result.Users {
		user := domain.User{
			ID:        v.Id,
			FirstName: v.Firstname,
			LastName:  v.Lastname,
			Username:  v.Username,
			CreatedAt: sql.NullTime{Time: v.CreatedAt.AsTime()},
		}

		users = append(users, user)
	}

	var res response.JsonRes
	res.Users = users

	c.JSON(http.StatusOK, res)
}

func (u userController) UploadImage(c *gin.Context) {
	f, err := c.FormFile("image")
	if err != nil {
		er.Handle(c, errors.New("bucket"))
		return
	}

	file, err := f.Open()
	if err != nil {

		er.Handle(c, err)
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		er.Handle(c, err)
		return
	}

	bucket, err := u.storage.DefaultBucket()
	if err != nil {
		er.Handle(c, err)
		return
	}
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	id := uuid.New()
	ext := filepath.Ext(f.Filename)
	obj := bucket.Object(fmt.Sprintf("pp/%s.%s", id.String(), ext))
	wc := obj.NewWriter(ctx)
	contentType := http.DetectContentType(buf)
	wc.ContentType = contentType
	if !strings.HasPrefix(contentType, "image/") {
		// handle error - file is not an image
		er.Handle(c, errors.New("unsupported image file"))
		return
	}

	if _, err := wc.Write(buf); err != nil {
		er.Handle(c, err)
		return
	}

	if err := wc.Close(); err != nil {
		er.Handle(c, err)
		return
	}

	attrs, err := obj.Attrs(ctx)
	if err != nil {
		er.Handle(c, fmt.Errorf("got error: %s", err))
		return
	}
	var res response.JsonRes
	res.Url = attrs.MediaLink
	c.JSON(http.StatusOK, res)

}

func (u userController) Close() error {
	return u.c.Close()
}
