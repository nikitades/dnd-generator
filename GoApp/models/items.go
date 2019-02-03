package models

import (
	"github.com/jmoiron/sqlx"
)

type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CategoryName string `json:"category_name" db:"category_name"`
	TypeName     string `json:"type_name" db:"type_name"`
	Icon         string `json:"icon"`
}

func GetRandomItemOfCategory(db *sqlx.DB, category ItemCategory) (Item, error) {
	item := Item{}
	err := db.Get(&item, `
		SELECT stuff.id, stuff.name, stuff_category.name as category_name, stuff_type.name as type_name, stuff.icon 
		FROM stuff
		LEFT JOIN stuff_category ON stuff_category.id = stuff.category_id
		LEFT JOIN stuff_type ON stuff_type.id = stuff_category.stuff_type_id
		WHERE stuff.category_id = $1
		ORDER BY RANDOM()
		LIMIT 1
	`, category.Id)
	if err != nil {
		return item, err
	}
	return item, nil
}
