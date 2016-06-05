package rest

import (
	"os"
	"encoding/json"
)

type RestDataSourceConfig struct {
	Name		string	`json:"name"`
	Type		string	`json:"type"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Host		string	`json:"host"`
	Port		int		`json:"port"`
	DbName		string	`json:"dbName"`
}

type RestBaseConfig struct {
	Host 				string	`json:"host"`
	LogPath 			string	`json:"logPath"`
	LogMaxSize 			int		`json:"logMaxSize"`
	LogMaxBackups 		int		`json:"logMaxBackups"`
	LogMaxAge 			int		`json:"logMaxAge"`
	DataSources			[]RestDataSourceConfig	`json:"datasources"`
}

func ReadConfiguration(path string) RestBaseConfig {
	file, err := os.Open(path)
	if err != nil {
		panic("Cannot open configuration file " + path)
	}

	decoder := json.NewDecoder(file)
	configuration := RestBaseConfig{}
	err = decoder.Decode(&configuration)

	if err != nil {
		panic("Cannot parse configuration file " + path)
	}


	return configuration
}
