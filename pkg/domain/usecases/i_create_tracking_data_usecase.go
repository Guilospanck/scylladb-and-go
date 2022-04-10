package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type ICreateTrackingDataUsecase interface {
	Perform(ctx context.Context, data dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)
}
