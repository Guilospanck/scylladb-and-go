package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type deleteTrackingDataByPartitionKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *deleteTrackingDataByPartitionKeyUsecase) Perform(ctx context.Context, data dtos.TrackingDataPartitionKeyDTO) error {
	return usecase.repo.DeleteTrackingDataByPartitionKey(ctx, &data)
}

func NewDeleteTrackingDataByPartitionKeyUsecase(repo interfaces.ITrackingDataRepository) *deleteTrackingDataByPartitionKeyUsecase {
	return &deleteTrackingDataByPartitionKeyUsecase{
		repo,
	}
}
