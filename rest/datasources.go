package rest

import (
	"github.com/op/go-logging"
	"time"
	"github.com/jmoiron/sqlx"
	"fmt"
)

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
