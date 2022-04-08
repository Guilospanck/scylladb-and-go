package interfaces

import "base/pkg/infrastructure/database/entities"

type ITrackingDataRepository interface {
	AddTrackingData(trackingData *entities.TrackingDataEntity) (*entities.TrackingDataEntity, error)
	DeleteTrackingDataByPrimaryKey(trackingData *entities.TrackingDataEntity) error
	DeleteTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) error
	FindTrackingDataByPrimaryKey(trackingData *entities.TrackingDataEntity) (*entities.TrackingDataEntity, error)
	FindAllTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) ([]entities.TrackingDataEntity, error)
	FindAllTrackingData() ([]entities.TrackingDataEntity, error)
}
