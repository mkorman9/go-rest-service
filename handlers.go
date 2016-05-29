package main

import (
	"net/http"
	"fmt"

	"github.com/mkorman9/restapi/rest"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cats := []Cat{ Cat{"Jack", 9}, Cat{"Daniels", 11} }
	rest.RespondJson(w, http.StatusOK, cats)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var cat Cat
	if err := rest.ReadJson(r, &cat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("%s, %d\n", cat.Name, cat.Age)
}
