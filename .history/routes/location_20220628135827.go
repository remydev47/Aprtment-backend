package route

import "github.com/kataras/iris/v12"

func Autocomplete(ctx iris.Context) {
	limit := "10"
	location := ctx.UrlParam("location")
}

func Search(ctx iris.Context) {
	
}