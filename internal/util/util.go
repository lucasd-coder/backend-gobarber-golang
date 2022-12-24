package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

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

func DateUtils(aux time.Time, layOut string) (*time.Time, error) {
	timeStampString := aux.Format(layOut)
	timeStamp, err := time.Parse(layOut, timeStampString)

	return &timeStamp, err
}

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsAfter(compareDate time.Time) bool {
	return time.Now().After(compareDate)
}

func ParseFromHttpResponse(resp *http.Response, model interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(bodyBytes, &model)

	if err != nil {
		return err
	}

	return nil
}
