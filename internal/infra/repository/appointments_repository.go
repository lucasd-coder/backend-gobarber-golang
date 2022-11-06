package repository

import (
	"fmt"
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/pkg/logger"

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
	err := db.Connection.Create(&appointment).Error
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
	var appointments []*model.Appointment
	parsedMonth, err := fmt.Printf("'%02d'", data.Month)
	if err != nil {
		logger.Log.Error(err.Error())
		return []*model.Appointment{}
	}

	dateFieldName := time.Date(data.Year, time.Month(parsedMonth), 0, 0, 0, 0, 0, time.UTC)
	date := dateFieldName.Format("2006-01-02")

	db.Connection.Model(model.Appointment{}).Where("provider_id = ? AND date LIKE ?", data.ProviderID, "%"+date+"%").Find(&appointments)
	return appointments
}

func (db *AppointmentsRepository) FindAllInDayFromProvider(data *dtos.FindAllInDayFromProviderDTO) []*model.Appointment {
	var appointments []*model.Appointment
	parsedMonth, err := fmt.Printf("'%02d'", data.Month)
	if err != nil {
		logger.Log.Error(err.Error())
		return []*model.Appointment{}
	}
	parsedDay, err := fmt.Printf("'%02d'", data.Day)
	if err != nil {
		logger.Log.Error(err.Error())
		return []*model.Appointment{}
	}

	dateFieldName := time.Date(data.Year, time.Month(parsedMonth), parsedDay, 0, 0, 0, 0, time.UTC)
	date := dateFieldName.Format("2006-01-02")

	db.Connection.Model(model.Appointment{}).Where("provider_id = ? AND date LIKE ?", data.ProviderID, "%"+date+"%").Find(&appointments)
	return appointments
}
