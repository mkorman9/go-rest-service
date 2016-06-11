package main

import (
	"github.com/mkorman9/restapi/rest"
	"github.com/jmoiron/sqlx"
)

func ReadAllCats(cats interface{}) {
	db := rest.GetContext().GetMember("db_Default").(*sqlx.DB)
	db.Select(cats, "SELECT * FROM CAT")
}

func AddCat(cat Cat) {
	db := rest.GetContext().GetMember("db_Default").(*sqlx.DB)

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO CAT (ROLE_NAME, NAME, DUELS_WON) VALUES (:ROLE_NAME, :NAME, :DUELS_WON)", &cat)
	tx.Commit()
}
