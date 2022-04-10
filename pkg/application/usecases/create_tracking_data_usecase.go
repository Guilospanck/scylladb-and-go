package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type createTrackingDataUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *createTrackingDataUsecase) Perform(ctx context.Context, data dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error) {
	return usecase.repo.AddTrackingData(ctx, &data)
}

func NewCreateTrackingDataUsecase(repo interfaces.ITrackingDataRepository) *createTrackingDataUsecase {
	return &createTrackingDataUsecase{
		repo,
	}
}
