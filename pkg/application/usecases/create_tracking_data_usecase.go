package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type createTrackingDataUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *createTrackingDataUsecase) Perform(data dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error) {
	return usecase.repo.AddTrackingData(&data)
}

func NewCreateTrackingDataUsecase(repo interfaces.ITrackingDataRepository) *createTrackingDataUsecase {
	return &createTrackingDataUsecase{
		repo,
	}
}
