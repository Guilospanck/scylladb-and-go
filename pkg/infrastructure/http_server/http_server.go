package httpserver

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	router *gin.Engine
}

func (server *HTTPServer) Setup() {
	server.router = gin.New()
}

func (server *HTTPServer) RegisterRoute(method string, endpoint string, handler ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		server.router.POST(endpoint, handler...)
	case "GET":
		server.router.GET(endpoint, handler...)
	case "PUT": // Put updates the entire resource
		server.router.PUT(endpoint, handler...)
	case "PATCH": // Patch updates partially
		server.router.PATCH(endpoint, handler...)
	case "DELETE":
		server.router.DELETE(endpoint, handler...)
	default:
		return errors.New("Method now allowed")
	}

	return nil
}
