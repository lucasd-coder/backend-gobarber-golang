package service

import (
	"backend-gobarber-golang/internal/dtos"
	"backend-gobarber-golang/internal/infra/repository"
	"backend-gobarber-golang/internal/infra/storage"
	"backend-gobarber-golang/internal/interfaces"
	"backend-gobarber-golang/internal/model"
)

type ListProviderAppointmentsService struct {
	AppointmentsRepository interfaces.AppointmentsRepository
	CacheProvider          interfaces.CacheProvider
}

func NewListProviderAppointmentsService(appointmentsRepository *repository.AppointmentsRepository, cacheProvider *storage.CacheProvider) *ListProviderAppointmentsService {
	return &ListProviderAppointmentsService{
		AppointmentsRepository: appointmentsRepository,
		CacheProvider:          cacheProvider,
	}
}

func (service *ListProviderAppointmentsService) Execute(dto *dtos.FindAllInDayFromProviderDTO) ([]model.Appointment, error) {
	return nil, nil
}
