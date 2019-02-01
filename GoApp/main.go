package main

import (
	"dnd-generator/GoApp/controllers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	var c = controllers.NewMainController()
	router := fasthttprouter.New()
	router.GET("/", c.MainPageHandler)
	router.GET("/getStartInfo", c.GetStartInfo)
	router.POST("/getItems", c.GetItems)
	router.ServeFiles("/static/*filepath", "../public_html")
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
