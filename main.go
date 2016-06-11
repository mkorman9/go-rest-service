package main

import (
	"net/http"
	"log"

	"github.com/mkorman9/restapi/rest"
)

func main() {
	configFileLocation := "rest_config.json"

	config := rest.ReadConfiguration(configFileLocation)
	context := rest.RestAppContext(AppRoutes, config)
	log.Fatal(http.ListenAndServe(config.Host, context))
}
