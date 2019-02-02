package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type ItemCategory struct {
	Id int `json:"id"`
}

func GetRandomCategoryOfTypes(db *sqlx.DB, typeIDs []int64) ItemCategory {
	category := ItemCategory{}
	var err error

	// TODO: Получить случайную категорию где типы такие как дано
	prequery, args, err := sqlx.In(`
		SELECT id
		FROM stuff_category
		WHERE stuff_type_id IN (?)
		ORDER BY RANDOM()
		LIMIT 1
	`, typeIDs)

	if err != nil {
		log.Println(err)
	}
	query := db.Rebind(prequery)
	err = db.Get(&category, query, args...)
	if err != nil {
		log.Println(err)
	}
	return category
}
