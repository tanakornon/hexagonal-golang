package kaleido

import "hex/internal/domain"

type Kaleido struct {
	apiKey string
}

func New(apiKey string) *Kaleido {
	return &Kaleido{
		apiKey: apiKey,
	}
}

func (o *Kaleido) GetWeatherData(city string) (*domain.WeatherData, error) {
	// Implement complex messy hard to understand logic

	return &domain.WeatherData{
		City:        city,
		Temperature: "100 celsius",
		Humidity:    "60%",
		Condition:   "so hot! from Kaleido",
	}, nil
}
