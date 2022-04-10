package interfaces

import (
	"base/pkg/domain/dtos"
	"context"
)

type ITrackingDataRepository interface {
	AddTrackingData(ctx context.Context, trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)

	DeleteTrackingDataByPrimaryKey(ctx context.Context, trackingData *dtos.TrackingDataPrimaryKeyDTO) error
	DeleteTrackingDataByPartitionKey(ctx context.Context, trackingData *dtos.TrackingDataPartitionKeyDTO) error

	FindTrackingDataByPrimaryKey(ctx context.Context, trackingData *dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error)
	FindAllTrackingDataByPartitionKey(ctx context.Context, trackingData *dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error)
	FindAllTrackingData(ctx context.Context) ([]*dtos.TrackingDataDTO, error)
}
