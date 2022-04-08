package interfaces

import (
	"base/pkg/domain/dtos"
)

type ITrackingDataRepository interface {
	AddTrackingData(trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)

	DeleteTrackingDataByPrimaryKey(trackingData *dtos.TrackingDataPrimaryKeyDTO) error
	DeleteTrackingDataByPartitionKey(trackingData *dtos.TrackingDataPartitionKeyDTO) error

	FindTrackingDataByPrimaryKey(trackingData *dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error)
	FindAllTrackingDataByPartitionKey(trackingData *dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error)
	FindAllTrackingData() ([]*dtos.TrackingDataDTO, error)
}
