/*
  Handlers are like Controllers
*/
package handlers

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/usecases"
)

type trackingDataHandler struct {
	logger interfaces.ILogger

	createUsecase usecases.ICreateTrackingDataUsecase

	deleteByPrimaryKey      usecases.IDeleteTrackingDataByPrimaryKeyUsecase
	deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase

	findByPrimaryKey      usecases.IFindTrackingDataByPrimaryKeyUsecase
	findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase
	findAll               usecases.IFindAllTrackingDataUsecase
}

func (handler *trackingDataHandler) Create() {}

func (handler *trackingDataHandler) DeleteByPrimaryKey() {}

func (handler *trackingDataHandler) DeleteAllByPartitionKey() {}

func (handler *trackingDataHandler) GetByPrimaryKey() {}

func (handler *trackingDataHandler) GetByAllByPartitionKey() {}

func (handler *trackingDataHandler) GetAll() {}

func NewTrackingDataHandler(
	logger interfaces.ILogger, createUsecase usecases.ICreateTrackingDataUsecase,
	deleteByPrimaryKey usecases.IDeleteTrackingDataByPrimaryKeyUsecase,
	deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase,
	findByPrimaryKey usecases.IFindTrackingDataByPrimaryKeyUsecase,
	findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase,
	findAll usecases.IFindAllTrackingDataUsecase,
) *trackingDataHandler {

	return &trackingDataHandler{
		logger,
		createUsecase,
		deleteByPrimaryKey,
		deleteAllByPartitionKey,
		findByPrimaryKey,
		findAllByPartitionKey,
		findAll,
	}
}
