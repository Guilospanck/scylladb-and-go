package usecases

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"context"
)

type findAllTrackingDataUsecase struct {
	repo interfaces.ITrackingDataRepository
}

func (usecase *findAllTrackingDataUsecase) Perform(ctx context.Context) ([]*dtos.TrackingDataDTO, error) {
	return usecase.repo.FindAllTrackingData(ctx)
}

func NewFindAllTrackingDataUsecase(repo interfaces.ITrackingDataRepository) *findAllTrackingDataUsecase {
	return &findAllTrackingDataUsecase{
		repo,
	}
}
