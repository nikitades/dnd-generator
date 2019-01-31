package main

import (
	"./controllers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	var a = controllers.MainController{}
	router := fasthttprouter.New()
	router.GET("/", a.Handle)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
