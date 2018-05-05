package service

import (
	"net/http"
	"sidekick/strader/logic/service/handlers/hokex"
	"sidekick/strader/logic/service/responses"
	"sidekick/strader/logic/service/svcerr"
	"xframe/handler/http_handler"
)

var (
	DEFAULT_PATH = "/strader"
)

var (
	httpHandlers = map[string]struct {
		Handler func(*http.Request) (interface{}, int)
	}{
		"/okex/future-trade":  {Handler: hokex.FutureTrade},
		"/okex/future-cancel": {Handler: hokex.FutureCancel},
	}
)

func init() {
	for route, task := range httpHandlers {
		http_handler.RegisterHTTPMuxHandler(DEFAULT_PATH+route, task.Handler)
	}
	http_handler.SUCCESS = svcerr.SUCCESS
	http_handler.DoBaseResponse = responses.DoBaseResponse
	http_handler.DoDataResponse = responses.DoDataResponse
}
