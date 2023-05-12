package main

import (
	"fmt"
	"hex/internal/adapters/inmemory"
	"hex/internal/adapters/openapi"

	"hex/internal/application"
)

func main() {
	apikey := "0x1213231231"

	api := openapi.NewOpenWeatherAPI(apikey)
	// k := kaleido.New(apikey)
	repo := inmemory.NewInMemoryWeatherRepository()

	app := application.NewWeatherUseCase(repo, api)

	app.UpdateWeatherData("Chonburi")
	data, _ := app.GetWeatherData("Bangkok")

	// router := echoclient.NewRouter(echoclient.UseCases{
	// 	WeatherUseCase: app,
	// })

	fmt.Printf("%v\n", data)
}
