package util

import (
	"bytes"
	"encoding/json"

	"backend-gobarber-golang/pkg/logger"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JsonLog(payload interface{}) *bytes.Buffer {
	objBody, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Info("Fail to parse payload log")
	}
	body := bytes.NewBuffer(objBody)
	return body
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
