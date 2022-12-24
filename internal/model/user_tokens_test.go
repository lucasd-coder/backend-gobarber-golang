package model_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestModel_NewUserToken(t *testing.T) {
	var id uuid.UUID = uuid.MustParse("0399e631-e2f0-4df5-b1d0-ca6d567a318c")

	userToken := model.NewUserToken(id.String())

	require.Equal(t, userToken.UserID, id.String())
}

func TestUserToken_BeforeCreate(t *testing.T) {
	gormDB, err := userTokenGormDB()
	assert.NoError(t, err)

	type fields struct {
		ID        uuid.UUID
		Token     uuid.UUID
		UserID    string
		User      model.User
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				ID:        uuid.MustParse("ae57a7e5-85ed-4f3b-9f14-1047d1e858a7"),
				Token:     uuid.MustParse("40d05c91-41e0-49d0-ae67-ea60b3bdbb31"),
				UserID:    "12f2c1a8-65b2-4f24-93a9-fe6aac6c4eea",
				CreatedAt: time.Date(2022, time.December, 20, 10, 0, 0, 0, time.UTC),
			},
			args: args{
				tx: gormDB,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userToken := &model.UserToken{
				ID:        tt.fields.ID,
				Token:     tt.fields.Token,
				UserID:    tt.fields.UserID,
				User:      tt.fields.User,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := userToken.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UserToken.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func userTokenGormDB() (*gorm.DB, error) {
	conn, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 conn,
		PreferSimpleProtocol: true,
	})

	return gorm.Open(dialector, &gorm.Config{})
}
