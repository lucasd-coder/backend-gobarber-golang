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

func TestModel_NewUser(t *testing.T) {
	name := "lucas"
	email := "lucas@gmail.com"
	password := "123456"

	user := model.NewUser(name, email, password)

	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	user.Email = ""
	require.NotEmpty(t, user)
}

func TestUser_BeforeCreate(t *testing.T) {
	gormDB, err := userGormDB()
	assert.NoError(t, err)

	type fields struct {
		ID        uuid.UUID
		Name      string
		Email     string
		Password  string
		Avatar    string
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
				ID: uuid.MustParse("ae57a7e5-85ed-4f3b-9f14-1047d1e858a7"),
			},
			args: args{
				tx: gormDB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &model.User{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				Avatar:    tt.fields.Avatar,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := user.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func userGormDB() (*gorm.DB, error) {
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
