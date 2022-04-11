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
	DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
	// GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse
}

type trackingDataHandler struct {
	logger              interfaces.ILogger
	httpResponseFactory factories.HttpResponseFactory

	createUsecase usecases.ICreateTrackingDataUsecase

	deleteByPrimaryKey      usecases.IDeleteTrackingDataByPrimaryKeyUsecase
	deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase

	findByPrimaryKey      usecases.IFindTrackingDataByPrimaryKeyUsecase
	findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase
	// findAll               usecases.IFindAllTrackingDataUsecase
}

func (handler *trackingDataHandler) Create(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataDTO{}

	/* Parse json and validate struct */
	err := dtos.ParseJson(httpRequest.Body, &dto, "TrackingDataDTO")
	if err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	/* Usecase */
	result, err := handler.createUsecase.Perform(httpRequest.Ctx, dto)
	if err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	return handler.httpResponseFactory.Created(result, nil)

}

func (handler *trackingDataHandler) DeleteByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataPrimaryKeyDTO{}

	/* Parse json and validate struct */
	if err := dtos.ParseJson(httpRequest.Body, &dto, "TrackingDataPrimaryKeyDTO"); err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	/* usecase */
	if err := handler.deleteByPrimaryKey.Perform(httpRequest.Ctx, dto); err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	return handler.httpResponseFactory.Ok(nil, nil)
}

func (handler *trackingDataHandler) DeleteAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataPartitionKeyDTO{}

	/* Parse json and validate struct */
	if err := dtos.ParseJson(httpRequest.Body, &dto, "TrackingDataPartitionKeyDTO"); err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	/* usecase */
	if err := handler.deleteAllByPartitionKey.Perform(httpRequest.Ctx, dto); err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	return handler.httpResponseFactory.Ok(nil, nil)
}

func (handler *trackingDataHandler) GetByPrimaryKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataPrimaryKeyDTO{}

	/* Parse json and validate struct */
	if err := dtos.ParseJson(httpRequest.Body, &dto, "TrackingDataPrimaryKeyDTO"); err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	/* usecase */
	result, err := handler.findByPrimaryKey.Perform(httpRequest.Ctx, dto)
	if err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	if result == nil {
		return handler.httpResponseFactory.NoContent(nil)
	}

	return handler.httpResponseFactory.Ok(result, nil)
}

func (handler *trackingDataHandler) GetByAllByPartitionKey(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
	dto := dtos.TrackingDataPartitionKeyDTO{}

	/* Parse json and validate struct */
	if err := dtos.ParseJson(httpRequest.Body, &dto, "TrackingDataPartitionKeyDTO"); err != nil {
		return handler.httpResponseFactory.BadRequest("Body must be a valid json.", nil)
	}

	/* usecase */
	result, err := handler.findAllByPartitionKey.Perform(httpRequest.Ctx, dto)
	if err != nil {
		return handler.httpResponseFactory.ErrorResponseMapper(err, nil)
	}

	if result == nil {
		return handler.httpResponseFactory.NoContent(nil)
	}

	return handler.httpResponseFactory.Ok(result, nil)
}

// func (handler *trackingDataHandler) GetAll(httpRequest httpserver.HttpRequest) httpserver.HttpResponse {
// }

func NewTrackingDataHandler(
	logger interfaces.ILogger, createUsecase usecases.ICreateTrackingDataUsecase,
	deleteByPrimaryKey usecases.IDeleteTrackingDataByPrimaryKeyUsecase,
	deleteAllByPartitionKey usecases.IDeleteTrackingDataByPartitionKeyUsecase,
	findByPrimaryKey usecases.IFindTrackingDataByPrimaryKeyUsecase,
	findAllByPartitionKey usecases.IFindAllTrackingDataByPartitionKeyUsecase,
	// findAll usecases.IFindAllTrackingDataUsecase,
) *trackingDataHandler {

	httpResponseFactory := factories.NewHttpResponseFactory()

	return &trackingDataHandler{
		logger,
		httpResponseFactory,
		createUsecase,
		deleteByPrimaryKey,
		deleteAllByPartitionKey,
		findByPrimaryKey,
		findAllByPartitionKey,
		// findAll,
	}
}
