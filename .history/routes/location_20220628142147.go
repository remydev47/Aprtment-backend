package route

import "github.com/kataras/iris/v12"

func Autocomplete(ctx iris.Context) {
	limit := "10"
	location := ctx.URLParam("location")
	limitQuery := ctx.URLParam("limit")
	if limitQuery != "" {
		limit = limitQuery
	}

	apiKey := os.GetEnv("LOCATION_TOKEN")
}

func Search(ctx iris.Context) {

}
