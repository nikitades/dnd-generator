package controllers

import (
	"dnd-generator/GoApp/models"
	"dnd-generator/GoApp/templates"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

const (
	MIN_ITEMS_COUNT = 1
	MAX_ITEMS_COUNT = 100
)

type MainController struct {
	db         *sqlx.DB
	bundleName string
}

type StartInfoResponse struct {
	Types []models.StuffType `json:"types"`
}

type ItemsResponse struct {
	AllTypes       []models.StuffType  `json:"types"`
	Count          int64               `json:"count"`
	Rarity         []models.ItemRarity `json:"rarity"`
	RequestedTypes []int64             `json:"requested_types"`
	Items          []models.Item       `json:"result"`
}

var defaultBundleName = "/bundles/app.min.js"

func NewMainController() *MainController {
	fmt.Println("Main controller created")
	mc := new(MainController)
	mc.bundleName = defaultBundleName
	var err error
	mc.db, err = sqlx.Connect("postgres", "user=postgres dbname=dnd password=dnd host=localhost sslmode=disable")
	if err != nil {
		log.Println(err)
		panic("Failed to connect to DB")
	}
	return mc
}

func (ctr *MainController) MainPageHandler(ctx *fasthttp.RequestCtx) {
	modTime, err := models.GetBundleMTime(ctr.bundleName)
	if err != nil {
		panic("Failed to get a bundle mod ")
	}
	markup := templates.Main(modTime)
	ctx.SetContentType("text/html; charset=utf8")
	fmt.Fprintf(ctx, markup)
}

func (ctr *MainController) GetStartInfo(ctx *fasthttp.RequestCtx) {
	types := models.GetStuffTypes(ctr.db)
	responseObj := StartInfoResponse{types}
	response, err := json.Marshal(responseObj)
	if err != nil {
		log.Println("Failed to serialize StartInfoResponse")
	}
	ctx.SetContentType("application/json; charset=utf8")
	ctx.SetBodyString(string(response))
}

func (ctr *MainController) GetItems(ctx *fasthttp.RequestCtx) {
	args := ctx.PostArgs()
	types := make([]int64, 0)
	var count int64
	var err error
	args.VisitAll(func(key, value []byte) {
		switch string(key) {
		case "types[]":
			var intTypeId int64
			intTypeId, err = strconv.ParseInt(string(value), 10, 0)
			if err != nil {
				log.Println("Wrong type ID given")
				break
			}
			types = append(types, intTypeId)
			break
		case "count":
			count, err = strconv.ParseInt(string(value), 10, 0)
			if err != nil {
				log.Println("Wrong items count given")
			}
			if count > MAX_ITEMS_COUNT {
				count = MAX_ITEMS_COUNT
			}
			if count < MIN_ITEMS_COUNT {
				count = MIN_ITEMS_COUNT
			}
		}
	})
	items := ctr.getItemsOfTypeAndCount(types, count)
	responseObj := ItemsResponse{
		AllTypes:       models.GetStuffTypes(ctr.db),
		Count:          count,
		Rarity:         models.GetRarities(ctr.db),
		RequestedTypes: types,
		Items:          items,
	}
	response, err := json.Marshal(responseObj)
	if err != nil {
		log.Println("Failed to serialize ItemsResponse")
	}
	ctx.SetContentType("application/json; charset=utf8")
	ctx.SetBodyString(string(response))
}

func (ctr *MainController) getItemsOfTypeAndCount(types []int64, count int64) []models.Item {
	outputItems := []models.Item{}
	for i := 0; i < int(count); i++ {
		randomCategory := models.GetRandomCategoryOfTypes(ctr.db, types)
		randomItemsOfCategory := models.GetRandomItemOfCategory(ctr.db, randomCategory)
		outputItems = append(outputItems, randomItemsOfCategory)
	}
	return outputItems
}
