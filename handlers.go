package main

import (
	"net/http"

	"github.com/mkorman9/restapi/rest"
	"github.com/jmoiron/sqlx"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var cats = []Cat{}
	db := rest.GetContext().GetMember("db_Default").(*sqlx.DB)
	db.Select(&cats, "SELECT * FROM CAT")

	rest.RespondJson(w, http.StatusOK, cats)
}

func Save(w http.ResponseWriter, r *http.Request) {
	var cat Cat
	if err := rest.ReadJson(r, &cat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
