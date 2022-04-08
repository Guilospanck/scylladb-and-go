package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type deleteTrackingDataByPartitionKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *deleteTrackingDataByPartitionKeyUsecase) Perform(data dtos.TrackingDataPartitionKeyDTO) error {
	return usecase.repo.DeleteTrackingDataByPartitionKey(&data)
}

func NewDeleteTrackingDataByPartitionKeyUsecase(repo interfaces.ITrackingDataRepository) *deleteTrackingDataByPartitionKeyUsecase {
	return &deleteTrackingDataByPartitionKeyUsecase{
		repo,
	}
}
