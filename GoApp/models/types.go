package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type StuffType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetStuffTypes(db *sqlx.DB) []StuffType {
	types := []StuffType{}
	err := db.Select(&types, "SELECT id, name FROM stuff_type")
	if err != nil {
		log.Println(err)
	}
	return types
}