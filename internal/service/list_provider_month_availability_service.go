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

type ListProviderMonthAvailabilityService struct {
	AppointmentsRepository interfaces.AppointmentsRepository
}

func NewListProviderMonthAvailabilityService(appointmentsRepository *repository.AppointmentsRepository) *ListProviderMonthAvailabilityService {
	return &ListProviderMonthAvailabilityService{
		AppointmentsRepository: appointmentsRepository,
	}
}

func (service *ListProviderMonthAvailabilityService) Execute(dto *dtos.FindAllInMonthFromProviderDTO) ([]*dtos.ResponseAllInMonthFromProviderDTO, error) {
	if !util.IsValidUUID(dto.ProviderID) {
		return []*dtos.ResponseAllInMonthFromProviderDTO{}, &errs.AppError{
			Message: "provider_id invalid.",
			Code:    400,
		}
	}

	appointments := service.AppointmentsRepository.FindAllInMonthFromProvider(dto)

	numberOfBaysInMonth := time.Date(dto.Year, time.Month(dto.Month-1), 0, 0, 0, 0, 0, time.UTC)

	getDaysInMonth := numberOfBaysInMonth.Day()

	eachDayArray := make([]int, getDaysInMonth)

	for i := 1; i < len(eachDayArray); i++ {
		eachDayArray[i] = i + 1
	}

	responseProviderDto := make([]*dtos.ResponseAllInMonthFromProviderDTO, len(appointments))

	for _, day := range eachDayArray {
		compareDate := time.Date(dto.Year, time.Month(dto.Month-1), day, 23, 59, 59, 0, time.UTC)

		responseProviderDto = append(responseProviderDto, dtos.NewResponseAllInMonthFromProviderDTO(day,
			(hasAppointmentInDay(appointments, day) && util.IsAfter(compareDate))))
	}

	return responseProviderDto, nil
}

func hasAppointmentInDay(appointments []*model.Appointment, day int) bool {
	for i := range appointments {
		if appointments[i].Date.Day() == day {
			return true
		}
	}
	return false
}
