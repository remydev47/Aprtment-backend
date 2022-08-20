package route

import (
	"os"

	"github.com/kataras/iris/v12"
)

func Autocomplete(ctx iris.Context) {
	limit := "10"
	location := ctx.URLParam("location")
	limitQuery := ctx.URLParam("limit")
	if limitQuery != "" {
		limit = limitQuery
	}

	apiKey := os.GetEnv("LOCATION_TOKEN")
	url :=
		"https://api.locationiq.com/v1/autocomplete.php?key=" + apiKey + "&q=" + location + "&limit=" + limit
}

func Search(ctx iris.Context) {

}
