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

func TestModel_NewAppointment(t *testing.T) {
	userID := uuid.NewString()
	providerID := uuid.NewString()
	date := time.Now()

	appointment := model.NewAppointment(userID, providerID, date)

	require.Equal(t, appointment.ProviderID, providerID)
	require.Equal(t, appointment.Date, date)
	require.Equal(t, appointment.UserID, userID)
}

func TestAppointment_BeforeCreate(t *testing.T) {
	gormDB, err := appointmentGormDB()
	assert.NoError(t, err)

	type fields struct {
		ID         uuid.UUID
		User       model.User
		UserID     string
		Provider   model.User
		ProviderID string
		Date       time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appointment := &model.Appointment{
				ID:         tt.fields.ID,
				User:       tt.fields.User,
				UserID:     tt.fields.UserID,
				Provider:   tt.fields.Provider,
				ProviderID: tt.fields.ProviderID,
				Date:       tt.fields.Date,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := appointment.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Appointment.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func appointmentGormDB() (*gorm.DB, error) {
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
