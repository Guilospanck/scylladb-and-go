package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type findAllTrackingDataByPartitionKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findAllTrackingDataByPartitionKeyUsecase) Perform(ctx context.Context, data dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindAllTrackingDataByPartitionKey(ctx, &data)
}

func NewFindAllTrackingDataByPartitionKeyUsecase(repo interfaces.ITrackingDataRepository) *findAllTrackingDataByPartitionKeyUsecase {
	return &findAllTrackingDataByPartitionKeyUsecase{
		repo,
	}
}
