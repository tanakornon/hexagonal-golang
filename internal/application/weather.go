package application

import "hex/internal/domain"

type WeatherUseCase interface {
	UpdateWeatherData(city string) error
	GetWeatherData(city string) (*domain.WeatherData, error)
}

type WeatherAPI interface {
	GetWeatherData(city string) (*domain.WeatherData, error)
}

type weatherUseCase struct {
	weatherRepo domain.WeatherRepository
	weatherAPI  WeatherAPI
}

func NewWeatherUseCase(weatherRepo domain.WeatherRepository, weatherAPI WeatherAPI) WeatherUseCase {
	return &weatherUseCase{
		weatherRepo: weatherRepo,
		weatherAPI:  weatherAPI,
	}
}

func (w *weatherUseCase) UpdateWeatherData(city string) error {
	weather, err := w.weatherAPI.GetWeatherData(city)
	if err != nil {
		return err
	}

	return w.weatherRepo.Save(weather)
}

func (w *weatherUseCase) GetWeatherData(city string) (*domain.WeatherData, error) {
	return w.weatherRepo.GetByCity(city)
}
