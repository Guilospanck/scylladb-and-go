package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type deleteTrackingDataByPrimaryKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *deleteTrackingDataByPrimaryKeyUsecase) Perform(ctx context.Context, data dtos.TrackingDataPrimaryKeyDTO) error {
	return usecase.repo.DeleteTrackingDataByPrimaryKey(ctx, &data)
}

func NewDeleteTrackingDataByPrimaryKeyUsecase(repo interfaces.ITrackingDataRepository) *deleteTrackingDataByPrimaryKeyUsecase {
	return &deleteTrackingDataByPrimaryKeyUsecase{
		repo,
	}
}
