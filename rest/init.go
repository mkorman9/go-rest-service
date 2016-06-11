package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/natefinch/lumberjack"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func RestAppContext(routes []Route, config RestBaseConfig) *mux.Router {
	GetContext().AddMember(RestConfig, config)

	initLogger(config)
	initDatasources(config)
	return initRouter(routes)
}

func initRouter(routes []Route) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler = route.HandlerFunc
		handler = AccessLogsHandler(handler, route.Name)
		handler = RecoverHandler(handler)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}
	return router
}

func initLogger(config RestBaseConfig) {
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
