package models

import (
	"github.com/jmoiron/sqlx"
)

type ItemRarity struct {
	Id   int
	Name string
	Code string
}

func GetRarities(db *sqlx.DB) ([]ItemRarity, error) {
	rarities := []ItemRarity{}
	err := db.Select(&rarities, "SELECT id, name, code FROM stuff_rarity")
	if err != nil {
		return rarities, err
	}
	return rarities, nil
}
