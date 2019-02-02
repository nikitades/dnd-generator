package models

import (
	"log"
	"github.com/jmoiron/sqlx"
)

type ItemRarity struct {
	Id   int
	Name string
	Code string
}

func GetRarities(db *sqlx.DB) []ItemRarity {
	rarities := []ItemRarity{}
	err := db.Select(&rarities, "SELECT id, name, code FROM stuff_rarity")
	if err != nil {
		log.Println(err)
	}
	return rarities
}
