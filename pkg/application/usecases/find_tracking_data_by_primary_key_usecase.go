package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type findTrackingDataByPrimaryKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findTrackingDataByPrimaryKeyUsecase) Perform(data dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindTrackingDataByPrimaryKey(&data)
}

func NewFindTrackingDataByPrimaryKeyUsecase(repo interfaces.ITrackingDataRepository) *findTrackingDataByPrimaryKeyUsecase {
	return &findTrackingDataByPrimaryKeyUsecase{
		repo,
	}
}
