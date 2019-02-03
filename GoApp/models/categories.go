package models

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ItemCategory struct {
	Id int `json:"id"`
}

func GetRandomCategoryOfTypes(db *sqlx.DB, typeIDs []int64) (ItemCategory, error) {
	category := ItemCategory{}
	var err error

	if len(typeIDs) == 0 {
		return category, errors.New("Category not found")
	}

	prequery, args, err := sqlx.In(`
		SELECT id
		FROM stuff_category
		WHERE stuff_type_id IN (?)
		ORDER BY RANDOM()
		LIMIT 1
	`, typeIDs)

	if err != nil {
		return category, err
	}
	query := db.Rebind(prequery)
	fmt.Println(query)
	err = db.Get(&category, query, args...)
	if err != nil {
		return category, err
	}
	return category, nil
}
