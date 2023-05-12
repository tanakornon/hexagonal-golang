package echoclient

import (
	"hex/internal/application"
)

type UseCases struct {
	WeatherUseCase application.WeatherUseCase
}

func NewRouter(ucs UseCases) interface{} {
	// weatherController := controller.NewWeatherController(ucs.weatherUseCase)

	// e := echo.New()
	return 0
}
