package service

import (
	"github.com/lucasd-coder/backend-gobarber-golang/internal/dtos"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/infra/repository"
	"github.com/lucasd-coder/backend-gobarber-golang/internal/interfaces"
)

type ListProviderDayAvailabilityService struct {
	AppointmentsRepository interfaces.AppointmentsRepository
}

func NewListProviderDayAvailabilityService(appointmentsRepository *repository.AppointmentsRepository) *ListProviderDayAvailabilityService {
	return &ListProviderDayAvailabilityService{
		AppointmentsRepository: appointmentsRepository,
	}
}

func (service *ListProviderDayAvailabilityService) Execute(dto *dtos.FindAllInDayFromProviderDTO) {}
