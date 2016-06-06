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

func initDatasources(config RestBaseConfig) {
	for _, datasourceSpec := range config.DataSources {
		datasource, err := connectDatasource(datasourceSpec)

		if err != nil {
			panic("Cannot open database connection for name " + datasourceSpec.Name + ", " + err.Error())
		}

		datasource.SetConnMaxLifetime(0)

		GetContext().AddMember("dbSpec_" + datasourceSpec.Name, datasourceSpec)
		GetContext().AddMember("db_" + datasourceSpec.Name, datasource)
	}

	go datasourcesWatchdog()
}

func datasourcesWatchdog() {
	var log = logging.MustGetLogger("db_watchdog_log")

	time.Sleep(1 * time.Minute)

	for _, datasourceName := range GetContext().GetDatasourcesList() {
		datasource := GetContext().GetMember(datasourceName).(*sqlx.DB)
		if datasource.Ping() != nil {
			config := GetContext().GetMember("dbSpec_" + datasourceName[3:]).(RestDataSourceConfig)
			log.Errorf("[Watchdog] Ping to datasource %s timed out, trying to reconnect", config.Name)

			datasource.Close()
			newDatasource, err := connectDatasource(config)

			if err != nil {
				log.Errorf("[Watchdog] Cannot reconnect to datasource %s, error %s", config.Name, err.Error())
			} else {
				log.Errorf("[Watchdog] Successfully reconnected to datasource %s", config.Name)
				GetContext().AddMember("db_" + config.Name, newDatasource)
			}
		}
	}
}

func connectDatasource(datasourceSpec RestDataSourceConfig) (*sqlx.DB, error) {
	return sqlx.Connect(datasourceSpec.Type, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		datasourceSpec.Username,
		datasourceSpec.Password,
		datasourceSpec.Host,
		datasourceSpec.Port,
		datasourceSpec.DbName))
}
