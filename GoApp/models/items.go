package models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type ItemCategory struct {
	Id int `json:"id"`
}

type Item struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CategoryName string `json:"category_name"`
	TypeName     string `json:"type_name"`
	Icon         string `json:"icon"`
}

func GetItems(db *sqlx.DB, typeIDs []int, count int) []Item {
	categories := []ItemCategory{}
	var err error
	/*
		TODO:
			1. Применить на первый запрос sqlx.In
			2. Ребинд
			3. Во втором запросе тоже ребинд и заюзать
	*/
	err = db.Select(&categories, `SELECT id 
	FROM stuff_category 
	WHERE stuff_type_id IN (?)
	ORDER BY RANDOM() 
	LIMIT 1`, typeIDs)
	items := []Item{}
	err = db.Select(&items, "SELECT")
	if err != nil {
		log.Println(err)
	}
	return items
}
