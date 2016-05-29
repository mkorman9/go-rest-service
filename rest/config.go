package rest

import (
	"os"
	"encoding/json"
)

type RestConfig struct {
	Host 			string	`json:"host"`
	LogPath 		string	`json:"logPath"`
	LogMaxSize 		int		`json:"logMaxSize"`
	LogMaxBackups 	int		`json:"logMaxBackups"`
	LogMaxAge 		int		`json:"logMaxAge"`
}

func ReadConfiguration(path string) RestConfig {
	file, err := os.Open(path)
	if err != nil {
		panic("Cannot open configuratuion file " + path)
	}

	decoder := json.NewDecoder(file)
	configuration := RestConfig{}
	err = decoder.Decode(&configuration)

	if err != nil {
		panic("Cannot parse configuratuion file " + path)
	}

	return configuration
}
