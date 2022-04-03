package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"backend-gobarber-golang/pkg/logger"

	"github.com/google/uuid"
)

type DiskStorageProvider struct{}

func NewDiskStorageProvider() *DiskStorageProvider {
	return &DiskStorageProvider{}
}

func (disk *DiskStorageProvider) SaveFile(file *multipart.FileHeader) string {
	ensureBaseDir("tmp/")

	openedFile, err := file.Open()
	if err != nil {
		logger.Log.Error(err.Error())
	}

	defer openedFile.Close()

	filename := fmt.Sprintf("%s-%s", uuid.NewString(), file.Filename)

	out, err := os.Create(filepath.Join("tmp/", filepath.Base(filename)))
	if err != nil {
		logger.Log.Error(err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, openedFile)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	err = out.Sync()
	if err != nil {
		logger.Log.Error(err.Error())
	}

	return filename
}

func (disk *DiskStorageProvider) DeleteFile(file string) {
	err := os.Remove(filepath.Join("tmp/", filepath.Base(file)))
	if err != nil {
		logger.Log.Error(err.Error())
	}
}

func ensureBaseDir(fpath string) error {
	baseDir := path.Dir(fpath)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0o755)
}
