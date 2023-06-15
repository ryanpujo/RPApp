package controller

import (
	"io"

	"github.com/gin-gonic/gin"
)

type ProductCrud interface {
	SimpleCrud
	io.Closer
}

type SimpleCrud interface {
	ByIdGetter
	Creator
	ByIdDeletion
	ByIdUpdate
	ManyGetter
}

type UserCrudCloser interface {
	SimpleCrud
	UploadImage(c *gin.Context)
	io.Closer
}

// Creator interface
type Creator interface {
	Create(c *gin.Context)
}

// Getter interface
type Getter interface {
	ByIdGetter
	ByUsernameGetter
	ManyGetter
	ByEmailGetter
}

type ByIdGetter interface {
	GetById(c *gin.Context)
}

type ByUsernameGetter interface {
	GetByUsername(c *gin.Context)
}

type ManyGetter interface {
	GetMany(c *gin.Context)
}

type ByEmailGetter interface {
	GetByEmail(c *gin.Context)
}

// deletion interface
type ByIdDeletion interface {
	DeleteById(c *gin.Context)
}

type ByUsernameDeletion interface {
	DeleteByUsername(c *gin.Context)
}

type ByEmailDeletion interface {
	DeleteByEmail(c *gin.Context)
}

type Deletion interface {
	ByEmailDeletion
	ByIdDeletion
	ByUsernameDeletion
}

// Update interface
type ByIdUpdate interface {
	UpdateById(c *gin.Context)
}

type ByUsernameUpdate interface {
	UpdateByUsername(c *gin.Context)
}

type ByEmailUpdate interface {
	UpdateByEmail(c *gin.Context)
}

type Update interface {
	ByEmailUpdate
	ByUsernameUpdate
	ByIdUpdate
}
