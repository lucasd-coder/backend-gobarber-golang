package storage_test

import (
	"io/ioutil"
	"mime/multipart"
	"net/textproto"
	"os"
	"testing"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskStorageSaveFileSuccessfully(t *testing.T) {
	tmpDir := t.TempDir()

	fileData := multipart.FileHeader{
		Filename: "test_disk_storage",
		Header:   textproto.MIMEHeader{},
		Size:     23,
	}

	f := mockOsCreate(tmpDir)

	testDiskStorageProvider := storage.DiskStorageProvider{Create: f}

	filename := testDiskStorageProvider.SaveFile(&fileData)

	assert.NotEmpty(t, filename)
}

func TestDiskStorageDeleteFileSuccessfully(t *testing.T) {
	tmpDir := t.TempDir()
	f := mockOsCreate(tmpDir)

	var fileData string = "test_disk_storage"
	testDiskStorageProvider := storage.DiskStorageProvider{Create: f}

	testDiskStorageProvider.DeleteFile(fileData)
}

func mockOsCreate(tmpDir string) func(name string) (*os.File, error) {
	return func(name string) (*os.File, error) {
		return ioutil.TempFile(tmpDir, "*")
	}
}
