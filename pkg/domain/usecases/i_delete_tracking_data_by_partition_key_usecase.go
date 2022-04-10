package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type IDeleteTrackingDataByPartitionKeyUsecase interface {
	Perform(ctx context.Context, data dtos.TrackingDataPartitionKeyDTO) error
}
