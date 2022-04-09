package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
)

type findAllTrackingDataByPartitionKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findAllTrackingDataByPartitionKeyUsecase) Perform(data dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindAllTrackingDataByPartitionKey(&data)
}

func NewFindAllTrackingDataByPartitionKeyUsecase(repo interfaces.ITrackingDataRepository) *findAllTrackingDataByPartitionKeyUsecase {
	return &findAllTrackingDataByPartitionKeyUsecase{
		repo,
	}
}
