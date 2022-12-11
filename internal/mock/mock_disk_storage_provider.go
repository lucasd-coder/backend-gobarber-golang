package mock

import (
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type MockDiskStorageProvider struct {
	mock.Mock
}

func (mock *MockDiskStorageProvider) SaveFile(file *multipart.FileHeader) string {
	return ""
}

func (mock *MockDiskStorageProvider) DeleteFile(file string) {
}
