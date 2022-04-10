package adapters

import (
	"base/pkg/application/interfaces"
	httpserver "base/pkg/infrastructure/http_server"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
  Responsible for reading the income request to a http request struct
  and returning the response as JSON
*/
func HandlerAdapter(handle func(httpserver.HttpRequest) httpserver.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* Read body */
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error(fmt.Sprintf("[HandlerAdapter] Error while trying to read request body: %s\n", err))
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		/* Read params */
		params := make(map[string]string)
		for _, param := range ctx.Params {
			params[param.Key] = param.Value
		}

		/* Put information into a http request */
		request := httpserver.HttpRequest{
			Body:    body,
			Headers: ctx.Request.Header,
			Params:  params,
			Query:   ctx.Request.URL.Query(),
			Ctx:     ctx.Request.Context(),
		}

		result := handle(request)

		ctx.JSON(result.StatusCode, result.Body)
	}
}
