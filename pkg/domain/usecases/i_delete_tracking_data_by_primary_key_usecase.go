package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type IDeleteTrackingDataByPrimaryKeyUsecase interface {
	Perform(ctx context.Context, data dtos.TrackingDataPrimaryKeyDTO) error
}
