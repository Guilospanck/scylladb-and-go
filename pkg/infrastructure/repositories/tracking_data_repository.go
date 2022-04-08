package repositories

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"base/pkg/infrastructure/database/entities"

	"go.uber.org/zap"
)

type trackingDataRepository struct {
	queryBuilder interfaces.IQueryBuilder[entities.TrackingDataEntity]
	logger       interfaces.ILogger
}

func (repo *trackingDataRepository) AddTrackingData(trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error) {
	err := repo.queryBuilder.Insert(trackingDataDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error("Could not add tracking data. Error: ", zap.Error(err))
		return nil, err
	}

	return trackingData, nil
}

func (repo *trackingDataRepository) DeleteTrackingDataByPrimaryKey(trackingData *dtos.TrackingDataPrimaryKeyDTO) error {
	err := repo.queryBuilder.Delete(trackingDataPrimaryKeyDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error("Could not delete tracking data using primary key. Error: ", zap.Error(err))
		return err
	}

	return nil
}

func (repo *trackingDataRepository) DeleteTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) error {
	err := repo.queryBuilder.DeleteAllFromPartitioningKey(trackingData)
	if err != nil {
		repo.logger.Error("Could not delete tracking data using partition key. Error: ", zap.Error(err))
		return err
	}

	return nil
}

func (repo *trackingDataRepository) FindTrackingDataByPrimaryKey(trackingData *entities.TrackingDataEntity) (*entities.TrackingDataEntity, error) {
	result, err := repo.queryBuilder.Get(trackingData)
	if err != nil {
		repo.logger.Error("Could not find tracking data by primary key. Error: ", zap.Error(err))
		return nil, err
	}

	return result, err
}

func (repo *trackingDataRepository) FindAllTrackingDataByPartitionKey(trackingData *entities.TrackingDataEntity) ([]entities.TrackingDataEntity, error) {
	results, err := repo.queryBuilder.Select(trackingData)
	if err != nil {
		repo.logger.Error("Could not find all tracking data by partition key. Error: ", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func (repo *trackingDataRepository) FindAllTrackingData() ([]entities.TrackingDataEntity, error) {
	results, err := repo.queryBuilder.SelectAll()
	if err != nil {
		repo.logger.Error("Could not find all tracking data. Error: ", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func trackingDataDTOToEntity(dto *dtos.TrackingDataDTO) *entities.TrackingDataEntity {
	trackingDataEntity := &entities.TrackingDataEntity{
		FirstName:       dto.FirstName,
		LastName:        dto.LastName,
		Timestamp:       dto.Timestamp,
		Location:        dto.Location,
		Speed:           dto.Speed,
		Heat:            dto.Heat,
		TelepathyPowers: dto.TelepathyPowers,
	}

	return trackingDataEntity
}

func trackingDataPrimaryKeyDTOToEntity(dto *dtos.TrackingDataPrimaryKeyDTO) *entities.TrackingDataEntity {
	trackingDataEntity := &entities.TrackingDataEntity{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Timestamp: dto.Timestamp,
	}

	return trackingDataEntity
}

func trackingDataPartitionKeyDTOToEntity(dto *dtos.TrackingDataPartitionKeyDTO) *entities.TrackingDataEntity {
	trackingDataEntity := &entities.TrackingDataEntity{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	return trackingDataEntity
}

func NewTrackingDataRepository(querybuilder interfaces.IQueryBuilder[entities.TrackingDataEntity], logger interfaces.ILogger) *trackingDataRepository {
	return &trackingDataRepository{
		queryBuilder: querybuilder,
		logger:       logger,
	}
}
