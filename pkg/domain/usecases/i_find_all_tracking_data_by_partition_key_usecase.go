package usecases

import (
	"base/pkg/domain/dtos"
	"context"
)

type IFindAllTrackingDataByPartitionKeyUsecase interface {
	Perform(ctx context.Context, data dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error)
}
