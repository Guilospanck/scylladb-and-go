package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type IFindTrackingDataByPrimaryKeyUsecase interface {
	Perform(ctx context.Context, data dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error)
}
