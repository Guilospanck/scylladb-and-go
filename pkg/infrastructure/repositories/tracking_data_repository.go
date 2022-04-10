package repositories

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"base/pkg/infrastructure/database/entities"
	"context"
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05 -0700 MST"

type trackingDataRepository struct {
	queryBuilder interfaces.IQueryBuilder[entities.TrackingDataEntity]
	logger       interfaces.ILogger
}

func (repo *trackingDataRepository) AddTrackingData(ctx context.Context, trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error) {
	err := repo.queryBuilder.Insert(ctx, trackingDataDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not add tracking data. Error:  %s", err.Error()))
		return nil, err
	}

	return trackingData, nil
}

func (repo *trackingDataRepository) DeleteTrackingDataByPrimaryKey(ctx context.Context, trackingData *dtos.TrackingDataPrimaryKeyDTO) error {
	err := repo.queryBuilder.Delete(ctx, trackingDataPrimaryKeyDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not delete tracking data using primary key. Error:  %s", err.Error()))
		return err
	}

	return nil
}

func (repo *trackingDataRepository) DeleteTrackingDataByPartitionKey(ctx context.Context, trackingData *dtos.TrackingDataPartitionKeyDTO) error {
	err := repo.queryBuilder.DeleteAllFromPartitioningKey(ctx, trackingDataPartitionKeyDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not delete tracking data using partition key. Error:  %s", err.Error()))
		return err
	}

	return nil
}

func (repo *trackingDataRepository) FindTrackingDataByPrimaryKey(ctx context.Context, trackingData *dtos.TrackingDataPrimaryKeyDTO) (*dtos.TrackingDataDTO, error) {
	result, err := repo.queryBuilder.Get(ctx, trackingDataPrimaryKeyDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not find tracking data by primary key. Error:  %s", err.Error()))
		return nil, err
	}

	return trackingDataEntityToDTO(result), err
}

func (repo *trackingDataRepository) FindAllTrackingDataByPartitionKey(ctx context.Context, trackingData *dtos.TrackingDataPartitionKeyDTO) ([]*dtos.TrackingDataDTO, error) {
	results, err := repo.queryBuilder.Select(ctx, trackingDataPartitionKeyDTOToEntity(trackingData))
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not find all tracking data by partition key. Error:  %s", err.Error()))
		return nil, err
	}

	var arrayOfResults []*dtos.TrackingDataDTO

	for _, value := range results {
		arrayOfResults = append(arrayOfResults, trackingDataEntityToDTO(&value))
	}

	return arrayOfResults, nil
}

func (repo *trackingDataRepository) FindAllTrackingData(ctx context.Context) ([]*dtos.TrackingDataDTO, error) {
	results, err := repo.queryBuilder.SelectAll(ctx)
	if err != nil {
		repo.logger.Error(fmt.Sprintf("Could not find all tracking data. Error:  %s", err.Error()))
		return nil, err
	}

	var arrayOfResults []*dtos.TrackingDataDTO

	for _, value := range results {
		arrayOfResults = append(arrayOfResults, trackingDataEntityToDTO(&value))
	}

	return arrayOfResults, nil
}

func trackingDataEntityToDTO(entity *entities.TrackingDataEntity) *dtos.TrackingDataDTO {
	return &dtos.TrackingDataDTO{
		FirstName:       entity.FirstName,
		LastName:        entity.LastName,
		Timestamp:       entity.Timestamp.String(),
		Location:        entity.Location,
		Speed:           entity.Speed,
		Heat:            entity.Heat,
		TelepathyPowers: entity.TelepathyPowers,
	}
}

func trackingDataDTOToEntity(dto *dtos.TrackingDataDTO) *entities.TrackingDataEntity {
	timestamp, _ := time.Parse(TIME_LAYOUT, dto.Timestamp)

	trackingDataEntity := &entities.TrackingDataEntity{
		FirstName:       dto.FirstName,
		LastName:        dto.LastName,
		Timestamp:       timestamp,
		Location:        dto.Location,
		Speed:           dto.Speed,
		Heat:            dto.Heat,
		TelepathyPowers: dto.TelepathyPowers,
	}

	return trackingDataEntity
}

func trackingDataPrimaryKeyDTOToEntity(dto *dtos.TrackingDataPrimaryKeyDTO) *entities.TrackingDataEntity {
	timestamp, _ := time.Parse(TIME_LAYOUT, dto.Timestamp)

	trackingDataEntity := &entities.TrackingDataEntity{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Timestamp: timestamp,
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
