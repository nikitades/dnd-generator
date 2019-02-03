package models

import (
	"github.com/jmoiron/sqlx"
)

type StuffType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetStuffTypes(db *sqlx.DB) ([]StuffType, error) {
	types := []StuffType{}
	err := db.Select(&types, "SELECT id, name FROM stuff_type")
	if err != nil {
		return nil, err
	}
	return types, nil
}
