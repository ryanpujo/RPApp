package mocks

import (
	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/mock"
)

type MockStorageClient struct {
	mock.Mock
}

func (m *MockStorageClient) DefaultBucket() (*storage.BucketHandle, error) {
	return &storage.BucketHandle{}, nil
}
