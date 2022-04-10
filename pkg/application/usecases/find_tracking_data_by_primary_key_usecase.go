package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type findTrackingDataByPrimaryKeyUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findTrackingDataByPrimaryKeyUsecase) Perform(ctx context.Context, data dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindTrackingDataByPrimaryKey(ctx, &data)
}

func NewFindTrackingDataByPrimaryKeyUsecase(repo interfaces.ITrackingDataRepository) *findTrackingDataByPrimaryKeyUsecase {
	return &findTrackingDataByPrimaryKeyUsecase{
		repo,
	}
}
