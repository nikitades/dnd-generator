package main

import (
	"dnd-generator/GoApp/controllers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	var c = controllers.MainController{}
	router := fasthttprouter.New()
	router.GET("/", c.Handle)
	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}
