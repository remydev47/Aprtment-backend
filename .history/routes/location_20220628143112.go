package route

import (
	"net/http"
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

	apiKey := os.Getenv("LOCATION_TOKEN")
	url :=
		"https://api.locationiq.com/v1/autocomplete.php?key=" + apiKey + "&q=" + location + "&limit=" + limit

	client := &http.Client{}
}

func Search(ctx iris.Context) {

}
