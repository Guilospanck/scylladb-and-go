package interfaces

import "base/pkg/infrastructure/database/entities"

type ITrackingDataRepository[T entities.TrackingDataEntity] interface {
	AddTrackingData(trackingData *T) (*T, error)
	DeleteTrackingDataByPrimaryKey(trackingData *T) error
	DeleteTrackingDataByPartitionKey(trackingData *T) error
	FindTrackingDataByPrimaryKey(trackingData *T) (*T, error)
	FindAllTrackingDataByPartitionKey(trackingData *T) ([]T, error)
	FindAllTrackingData() ([]T, error)
}
