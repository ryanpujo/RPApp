package registry

import (
	"context"

	"firebase.google.com/go/storage"
	"github.com/spriigan/broker/firebaseapp"
	"github.com/spriigan/broker/interface/controller"
	c "github.com/spriigan/broker/user/controller"
	"github.com/spriigan/broker/user/grpc/client"
)

func (r registry) NewUserClient() client.UserClientCloser {
	return client.NewUserClient()
}

func (r registry) NewUserController() controller.UserCrudCloser {
	return c.NewUserController(r.NewUserClient(), r.NewFirebaseStorage())
}

func (r registry) NewFirebaseStorage() *storage.Client {
	app := firebaseapp.New()
	client, err := app.Storage(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
