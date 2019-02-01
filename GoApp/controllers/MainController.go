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
)

type MainController struct {
	db         *sqlx.DB
	bundleName string
}

type StartInfoResponse struct {
	Types []models.StuffType `json:"types"`
}

type ItemsResponse struct {
	Items []models.Item `json:"items"`
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
	args := ctx.QueryArgs()
	fmt.Println(args)
	items := models.GetItems(ctr.db)
	responseObj := ItemsResponse{items}
	response, err := json.Marshal(responseObj)
	if err != nil {
		log.Println("Failed to serialize ItemsResponse")
	}
	ctx.SetContentType("application/json; charset=utf8")
	ctx.SetBodyString(string(response))
}
