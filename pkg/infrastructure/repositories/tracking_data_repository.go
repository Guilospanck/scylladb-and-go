package repositories

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database/entities"

	"go.uber.org/zap"
)

type trackingDataRepository[T entities.TrackingDataEntity] struct {
	queryBuilder interfaces.IQueryBuilder[T]
	logger       interfaces.ILogger
}

func (repo *trackingDataRepository[T]) AddTrackingData(trackingData *T) (*T, error) {
	err := repo.queryBuilder.Insert(trackingData)
	if err != nil {
		repo.logger.Error("Could not add tracking data. Error: ", zap.Error(err))
		return nil, err
	}

	return trackingData, nil
}

func (repo *trackingDataRepository[T]) DeleteTrackingDataByPrimaryKey(trackingData *T) error {
	err := repo.queryBuilder.Delete(trackingData)
	if err != nil {
		repo.logger.Error("Could not delete tracking data using primary key. Error: ", zap.Error(err))
		return err
	}

	return nil
}

func (repo *trackingDataRepository[T]) DeleteTrackingDataByPartitionKey(trackingData *T) error {
	err := repo.queryBuilder.DeleteAllFromPartitioningKey(trackingData)
	if err != nil {
		repo.logger.Error("Could not delete tracking data using partition key. Error: ", zap.Error(err))
		return err
	}

	return nil
}

func (repo *trackingDataRepository[T]) FindTrackingDataByPrimaryKey(trackingData *T) (*T, error) {
	result, err := repo.queryBuilder.Get(trackingData)
	if err != nil {
		repo.logger.Error("Could not find tracking data by primary key. Error: ", zap.Error(err))
		return nil, err
	}

	return result, err
}

func (repo *trackingDataRepository[T]) FindAllTrackingDataByPartitionKey(trackingData *T) ([]T, error) {
	results, err := repo.queryBuilder.Select(trackingData)
	if err != nil {
		repo.logger.Error("Could not find all tracking data by partition key. Error: ", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func (repo *trackingDataRepository[T]) FindAllTrackingData() ([]T, error) {
	results, err := repo.queryBuilder.SelectAll()
	if err != nil {
		repo.logger.Error("Could not find all tracking data. Error: ", zap.Error(err))
		return nil, err
	}

	return results, nil
}

func NewTrackingDataRepository[T entities.TrackingDataEntity](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger) *trackingDataRepository[T] {

	return &trackingDataRepository[T]{
		queryBuilder: querybuilder,
		logger:       logger,
	}
}
