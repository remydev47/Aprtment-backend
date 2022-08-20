package main 

import (
    "github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()

	   app := iris.Default()
	   
}