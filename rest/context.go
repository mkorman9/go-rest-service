package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/natefinch/lumberjack"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func RestAppContext(routes []Route, config RestConfig) *mux.Router {
	initLogger(config)
	return initRouter(routes)
}

func initRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = AccessLogsHandler(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}
	return router
}

func initLogger(config RestConfig) {
	logsBackend := logging.NewLogBackend(
		&lumberjack.Logger{
			Filename:   config.LogPath,
			MaxSize:    config.LogMaxSize,
			MaxBackups: config.LogMaxBackups,
			MaxAge:     config.LogMaxAge,
		},
		"", 0)

	var format = logging.MustStringFormatter(
		`[%{time:02.01.2006 15:04:05} %{level}] %{message}`,
	)
	logsBackendFormatter := logging.NewBackendFormatter(logsBackend, format)
	logging.SetBackend(logsBackendFormatter)
}
