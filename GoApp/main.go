package main

import (
	"dnd-generator/GoApp/controllers"
	"dnd-generator/GoApp/models"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	settings := models.LoadSettings("../env.ini")
	var c = controllers.NewMainController(settings)
	router := fasthttprouter.New()
	router.GET("/", c.MainPageHandler)
	router.GET("/getStartInfo", c.GetStartInfo)
	router.POST("/getItems", c.GetItems)
	router.ServeFiles("/static/*filepath", "../public_html")
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
