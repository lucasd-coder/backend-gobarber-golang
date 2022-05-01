package repository

import (
	"time"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/pkg/logger"

	"gorm.io/gorm"
)

type AppointmentsRepository struct {
	Connection *gorm.DB
}

func NewAppointmentsRepository(connectionDb *gorm.DB) *AppointmentsRepository {
	return &AppointmentsRepository{
		Connection: connectionDb,
	}
}

func (db *AppointmentsRepository) Save(appointment *model.Appointment) {
	err := db.Connection.Save(&appointment).Error
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
}

func (db *AppointmentsRepository) FindByDate(date *time.Time, providerId string) *model.Appointment {
	var appointment model.Appointment
	db.Connection.Model(model.Appointment{}).Where("date = ? AND provider_id = ?", date, providerId).Find(&appointment)
	return &appointment
}

func (db *AppointmentsRepository) FindAllInMonthFromProvider(data *dtos.FindAllInMonthFromProviderDTO) []*model.Appointment {
	return nil
}

func (db *AppointmentsRepository) FindAllInDayFromProvider(data *dtos.FindAllInDayFromProviderDTO) []*model.Appointment {
	return nil
}
