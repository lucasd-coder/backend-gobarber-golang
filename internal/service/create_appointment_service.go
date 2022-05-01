package service

import (
	"context"
	"fmt"
	"time"

	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/errs"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/infra/storage"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/model"
	"backend-gobarber-golang/internal/util"
	"backend-gobarber-golang/pkg/logger"

	"github.com/google/uuid"
)

type CreateAppointmentService struct {
	AppointmentsRepository  interfaces.AppointmentsRepository
	NotificationsRepository interfaces.NotificationsRepository
	CacheProvider           interfaces.CacheProvider
}

func NewCreateAppointmentService(appointmentsRepository *repository.AppointmentsRepository, notificationsRepository *repository.NotificationsRepository,
	cacheProvider *storage.CacheProvider,
) *CreateAppointmentService {
	return &CreateAppointmentService{
		AppointmentsRepository:  appointmentsRepository,
		NotificationsRepository: notificationsRepository,
		CacheProvider:           cacheProvider,
	}
}

func (service *CreateAppointmentService) Execute(userId string, dto *dtos.AppointmentDTO) (*model.Appointment, error) {
	if !util.IsValidUUID(dto.ProviderID) {
		return &model.Appointment{}, &errs.AppError{
			Message: "provider_id invalid.",
			Code:    400,
		}
	}

	appointmentDate := time.Now()

	timeStamp, err := util.DateUtils(dto.Date, "2006-01-02 15:04:05")
	if err != nil {
		return &model.Appointment{}, &errs.AppError{
			Message: "Invalid format date accepted format 2006-01-02 15:04:05",
			Code:    400,
		}
	}

	if appointmentDate.Before(*timeStamp) {
		return &model.Appointment{}, &errs.AppError{
			Message: "You can't create an appointment on a post date.",
			Code:    400,
		}
	}

	if userId == dto.ProviderID {
		return &model.Appointment{}, &errs.AppError{
			Message: "You can't create an appointment with yourself.",
			Code:    400,
		}
	}

	hr, _, _ := timeStamp.Clock()

	if hr < 8 || hr > 17 {
		return &model.Appointment{}, &errs.AppError{
			Message: "You can only create appointments between 8am and 5pm",
			Code:    400,
		}
	}

	date, _ := util.DateUtils(*timeStamp, "2006-01-02 15:04")

	findAppointnentInSameDate := service.AppointmentsRepository.FindByDate(date, dto.ProviderID)

	if findAppointnentInSameDate.ProviderID != "" {
		return &model.Appointment{}, &errs.AppError{
			Message: "This appointment is already backed",
			Code:    400,
		}
	}

	appointment := model.NewAppointment(userId, dto.ProviderID, *date)

	service.AppointmentsRepository.Save(appointment)

	content := fmt.Sprintf("Novo agendamento para dia %s",
		util.DateFormat(*timeStamp, "02-Jan-2006 15:04:05"))

	notification := model.NewNotification(uuid.MustParse(dto.ProviderID), content)

	err = service.NotificationsRepository.Save(notification)
	if err != nil {
		logger.Log.Error(err.Error())
		return &model.Appointment{}, &errs.AppError{
			Message: "Internal Server Error",
			Code:    500,
		}
	}

	key := fmt.Sprintf("provider-appointments:%s:%s", dto.ProviderID, util.DateFormat(*timeStamp, "2006-01-02"))

	service.CacheProvider.Invalidate(context.TODO(), key)

	return appointment, nil
}
