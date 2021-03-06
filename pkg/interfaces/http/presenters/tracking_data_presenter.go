/*
  Presenters are Routes
  Presenters will show the routes for some endpoint (controller/handler)
*/
package presenters

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/adapters"
	httpserver "base/pkg/infrastructure/http_server"
	"base/pkg/interfaces/http/handlers"
)

type trackingDataPresenter struct {
	logger  interfaces.ILogger
	handler handlers.ITrackingDataHandler
}

func (presenter *trackingDataPresenter) Register(httpServer httpserver.IHTTPServer) {
	httpServer.RegisterRoute("POST", "/api/v1/tracking", adapters.HandlerAdapter(presenter.handler.Create, presenter.logger))
	httpServer.RegisterRoute("DELETE", "/api/v1/tracking", adapters.HandlerAdapter(presenter.handler.DeleteByPrimaryKey, presenter.logger))
	httpServer.RegisterRoute("DELETE", "/api/v1/tracking/bypartition", adapters.HandlerAdapter(presenter.handler.DeleteAllByPartitionKey, presenter.logger))
	httpServer.RegisterRoute("GET", "/api/v1/tracking", adapters.HandlerAdapter(presenter.handler.GetByPrimaryKey, presenter.logger))
	httpServer.RegisterRoute("GET", "/api/v1/tracking/bypartition", adapters.HandlerAdapter(presenter.handler.GetByAllByPartitionKey, presenter.logger))
	httpServer.RegisterRoute("GET", "/api/v1/tracking/all", adapters.HandlerAdapter(presenter.handler.GetAll, presenter.logger))
}

func NewTrackingDataPresenters(logger interfaces.ILogger, handler handlers.ITrackingDataHandler) *trackingDataPresenter {
	return &trackingDataPresenter{
		logger,
		handler,
	}
}
