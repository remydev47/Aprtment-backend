package route

import "github.com/kataras/iris/v12"

func Autocomplete(ctx iris.Context) {
	limit := "10"
	location := ctx.URLParam("location")
	limitQuery := ctx.URLParam("limit")
}


func Search(ctx iris.Context) {
	
}