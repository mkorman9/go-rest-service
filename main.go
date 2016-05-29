package main

import (
	"os"
	"net/http"
	"log"

	"github.com/mkorman9/restapi/rest"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Application requires passing configuration file path as argument")
	}

	config := rest.ReadConfiguration(args[0])
	context := rest.RestAppContext(routes, config)
	log.Fatal(http.ListenAndServe(config.Host, context))
}
