package interfaces

import (
	"base/pkg/domain/dtos"
	"base/pkg/infrastructure/database/entities"
)

type ITrackingDataRepository interface {
	AddTrackingData(trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)
	DeleteTrackingDataByPrimaryKey(trackingData *dtos.TrackingDataPrimaryKeyDTO) error
	DeleteTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) error
	FindTrackingDataByPrimaryKey(trackingData *entities.TrackingDataEntity) (*entities.TrackingDataEntity, error)
	FindAllTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) ([]entities.TrackingDataEntity, error)
	FindAllTrackingData() ([]entities.TrackingDataEntity, error)
}
