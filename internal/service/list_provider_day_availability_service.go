package service

import (
	"time"

	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/errs"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/model"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/util"
)

type ListProviderDayAvailabilityService struct {
	AppointmentsRepository interfaces.AppointmentsRepository
}

func NewListProviderDayAvailabilityService(appointmentsRepository *repository.AppointmentsRepository) *ListProviderDayAvailabilityService {
	return &ListProviderDayAvailabilityService{
		AppointmentsRepository: appointmentsRepository,
	}
}

func (service *ListProviderDayAvailabilityService) Execute(dto *dtos.FindAllInDayFromProviderDTO) ([]*dtos.ResponseProviderDTO, error) {
	if !util.IsValidUUID(dto.ProviderID) {
		return []*dtos.ResponseProviderDTO{}, &errs.AppError{
			Message: "provider_id invalid.",
			Code:    400,
		}
	}

	appointments := service.AppointmentsRepository.FindAllInDayFromProvider(dto)

	var eachHourArray [10]int

	for i := 1; i < len(eachHourArray); i++ {
		eachHourArray[i] = i + 8
	}

	responseProviderDto := make([]*dtos.ResponseProviderDTO, len(appointments))

	for _, hour := range eachHourArray {
		compareDate := time.Date(dto.Year, time.Month(dto.Month-1), dto.Day, hour, 0, 0, 0, time.UTC)

		responseProviderDto = append(responseProviderDto,
			dtos.NewResponseProviderDTO(hour, (hasAppointmentInHour(appointments, hour) && isAfter(compareDate))))
	}

	return responseProviderDto, nil
}

func hasAppointmentInHour(appointments []*model.Appointment, hour int) bool {
	for i := range appointments {
		if appointments[i].Date.Hour() == hour {
			return true
		}
	}
	return false
}

func isAfter(compareDate time.Time) bool {
	return time.Now().After(compareDate)
}
