package httpserver

import (
	"context"
	"net/http"
)

type HttpRequest struct {
	Body    []byte
	Headers http.Header
	Params  map[string]string
	Query   map[string][]string
	Auth    interface{}
	Ctx     context.Context
}

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}
