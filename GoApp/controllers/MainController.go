package controllers

import (
	"github.com/valyala/fasthttp"
)

type MainController struct {
}

func (ctr *MainController) Handle(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Govno")
}
