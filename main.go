package main

import "github.com/mariotj/survey-app/api"

func main() {
	engine := api.Run()
	_ = engine.Run(":7099")
}
