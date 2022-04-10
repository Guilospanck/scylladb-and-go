package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type IFindAllTrackingDataUsecase interface {
	Perform(ctx context.Context) ([]*dtos.TrackingDataDTO, error)
}
