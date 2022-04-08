package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type deleteTrackingDataByPrimaryKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *deleteTrackingDataByPrimaryKeyUsecase) Perform(data dtos.TrackingDataPrimaryKeyDTO) error {
	return usecase.repo.DeleteTrackingDataByPrimaryKey(&data)
}

func NewDeleteTrackingDataByPrimaryKeyUsecase(repo interfaces.ITrackingDataRepository) *deleteTrackingDataByPrimaryKeyUsecase {
	return &deleteTrackingDataByPrimaryKeyUsecase{
		repo,
	}
}
