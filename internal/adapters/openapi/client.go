package openapi

import "hex/internal/domain"

type OpenWeatherAPI struct {
	apiKey string
}

func NewOpenWeatherAPI(apiKey string) *OpenWeatherAPI {
	return &OpenWeatherAPI{
		apiKey: apiKey,
	}
}

func (o *OpenWeatherAPI) GetWeatherData(city string) (*domain.WeatherData, error) {
	// Make API call and parse response

	return &domain.WeatherData{
		City:        city,
		Temperature: "100 celsius",
		Humidity:    "60%",
		Condition:   "so hot!",
	}, nil
}
