package presenters

import httpserver "base/pkg/infrastructure/http_server"

type IRoutes interface {
	Register(httpServer httpserver.IHTTPServer)
}
