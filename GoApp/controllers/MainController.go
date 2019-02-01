package controllers

import (
	"dnd-generator/GoApp/templates"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strconv"
)

type MainController struct {
	db         string
	bundleName string
}

var defaultBundleName = "/bundles/app.min.js"

func NewMainController() *MainController {
	mc := new(MainController)
	mc.bundleName = defaultBundleName
	mc.db = "le db connection"
	return mc
}

func (ctr *MainController) MainPageHandler(ctx *fasthttp.RequestCtx) {
	modTime, err := ctr.getBundleMTime()
	fmt.Println(modTime)
	if err != nil {
		panic("Failed to get a bundle mod ")
	}
	markup := templates.Main(modTime)
	ctx.SetContentType("text/html; charset=utf8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, markup)
}

func (ctr *MainController) getBundleMTime() (string, error) {
	stat, err := os.Stat("../public_html/" + ctr.bundleName)
	if err != nil {
		return "", err
	}
	bundleModTime := stat.ModTime().Unix()
	return strconv.FormatInt(bundleModTime, 10), nil
}
