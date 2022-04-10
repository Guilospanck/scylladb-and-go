/*
  Handlers are like Controllers
*/
package handlers

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/usecases"
	httpserver "base/pkg/infrastructure/http_server"
)

type ITrackingDataHandler interface {
	Create(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
}

type trackingDataHandler struct {
	logger interfaces.ILogger

	createUsecase usecases.ICreateTrackingDataUsecase

	deleteByPrimaryKey      usecases.IDeleteTrackingDataByPrimaryKeyUsecase
	deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase

	findByPrimaryKey      usecases.IFindTrackingDataByPrimaryKeyUsecase
	findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase
	findAll               usecases.IFindAllTrackingDataUsecase
}

func (handler *trackingDataHandler) Create(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

func (handler *trackingDataHandler) DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

func (handler *trackingDataHandler) DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

func (handler *trackingDataHandler) GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

func (handler *trackingDataHandler) GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

func (handler *trackingDataHandler) GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
}

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
