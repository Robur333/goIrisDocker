package main

import (
	"log"
	"restapi/internal"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	err:= godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }
	app := internal.IrisMainApp()
    // Listen and serve on http://localhost:8080.
    app.Run(iris.Addr(":8080"))
}
