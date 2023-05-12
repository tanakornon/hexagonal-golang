package controller

import "hex/internal/application"

type WeatherController struct {
	uc application.WeatherUseCase
}

func NewWeatherController(uc application.WeatherUseCase) *WeatherController {
	return &WeatherController{
		uc: uc,
	}
}

func (c *WeatherController) GetWeather() {

}
