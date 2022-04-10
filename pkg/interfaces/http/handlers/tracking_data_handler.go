/*
  Handlers are like Controllers
*/
package handlers

import (
	"base/pkg/application/interfaces"
	"base/pkg/domain/dtos"
	"base/pkg/domain/usecases"
	httpserver "base/pkg/infrastructure/http_server"
	"base/pkg/interfaces/http/factories"
)

type ITrackingDataHandler interface {
	Create(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
}

type trackingDataHandler struct {
	logger              interfaces.ILogger
	httpResponseFactory factories.HttpResponseFactory

	createUsecase usecases.ICreateTrackingDataUsecase

	// deleteByPrimaryKey      usecases.IDeleteTrackingDataByPrimaryKeyUsecase
	// deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase

	// findByPrimaryKey      usecases.IFindTrackingDataByPrimaryKeyUsecase
	// findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase
	// findAll               usecases.IFindAllTrackingDataUsecase
}

func (handler *trackingDataHandler) Create(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataDTO{}

	err := dtos.ParseJson(httpRequest.Body, dto, "TrackingDataDTO")
	if err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	results, err := handler.createUsecase.Perform(httpRequest.Ctx, dto)
	if err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	return handler.httpResponseFactory.Created(results, nil)

}

// func (handler *trackingDataHandler) DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

// func (handler *trackingDataHandler) DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

// func (handler *trackingDataHandler) GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

// func (handler *trackingDataHandler) GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

// func (handler *trackingDataHandler) GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

func NewTrackingDataHandler(
	logger interfaces.ILogger, createUsecase usecases.ICreateTrackingDataUsecase,
	// deleteByPrimaryKey usecases.IDeleteTrackingDataByPrimaryKeyUsecase,
	// deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase,
	// findByPrimaryKey usecases.IFindTrackingDataByPrimaryKeyUsecase,
	// findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase,
	// findAll usecases.IFindAllTrackingDataUsecase,
) *trackingDataHandler {

	httpResponseFactory := factories.NewHttpResponseFactory()

	return &trackingDataHandler{
		logger,
		httpResponseFactory,
		createUsecase,
		// deleteByPrimaryKey,
		// deleteAllByPartitionKey,
		// findByPrimaryKey,
		// findAllByPartitionKey,
		// findAll,
	}
}
