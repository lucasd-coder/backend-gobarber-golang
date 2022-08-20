package service

import (
	"context"
	"encoding/json"
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

	"github.com/go-redis/redis/v8"
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

func (service *ListProviderAppointmentsService) Execute(dto *dtos.FindAllInDayFromProviderDTO) ([]*model.Appointment, error) {
	if !util.IsValidUUID(dto.ProviderID) {
		return []*model.Appointment{}, &errs.AppError{
			Message: "provider_id invalid.",
			Code:    400,
		}
	}

	var existsProvider []*model.Appointment
	cacheKey := fmt.Sprintf("provider-appointments:%s:%d-%d-%d", dto.ProviderID, dto.Year, dto.Month, dto.Day)

	resultCache, err := service.CacheProvider.Recover(context.TODO(), cacheKey)
	if err != redis.Nil {
		logger.Log.Warnf("cache with key: %s not exists", cacheKey)
	}

	err = json.Unmarshal([]byte(resultCache), &existsProvider)
	if err != nil {
		logger.Log.Error(err)
	}

	if len(existsProvider) == 0 {
		existsProvider = service.AppointmentsRepository.FindAllInDayFromProvider(dto)

		err := service.CacheProvider.Save(context.TODO(), cacheKey, existsProvider, 4*time.Hour)
		if err != nil {
			logger.Log.Error(err)
		}

		return existsProvider, nil
	}

	return existsProvider, nil
}
