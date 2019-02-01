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
	router.ServeFiles("/static/*filepath", "../public_html")
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
