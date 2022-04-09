package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type findAllTrackingDataUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findAllTrackingDataUsecase) Perform() ([]*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindAllTrackingData()
}

func NewFindAllTrackingDataUsecase(repo interfaces.ITrackingDataRepository) *findAllTrackingDataUsecase {
	return &findAllTrackingDataUsecase{
		repo,
	}
}
