package main

import (
	"github.com/mkorman9/restapi/rest"
)

var routes = []rest.Route{
	rest.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	rest.Route{
		"Save",
		"POST",
		"/save",
		Save,
	},
}
