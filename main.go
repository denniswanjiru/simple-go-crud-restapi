package main

import "github.com/denniswanjiru/simple-go-crud-restapi/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":8080")
}
