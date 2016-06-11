package main

import (
	"net/http"

	"github.com/mkorman9/restapi/rest"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cats := []Cat{}
	ReadAllCats(&cats)
	rest.RespondJson(w, http.StatusOK, cats)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var cat Cat
	if err := rest.ReadJson(r, &cat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	AddCat(cat)
	rest.RespondJson(w, http.StatusOK, "ok")
}
