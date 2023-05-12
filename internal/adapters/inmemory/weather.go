package inmemory

import (
	"hex/internal/domain"
)

type InMemoryWeatherRepository struct {
	collection map[string]*domain.WeatherData
}

func NewInMemoryWeatherRepository() *InMemoryWeatherRepository {
	return &InMemoryWeatherRepository{
		collection: make(map[string]*domain.WeatherData, 0),
	}
}

func (m *InMemoryWeatherRepository) Save(weatherData *domain.WeatherData) error {
	m.collection[weatherData.City] = weatherData
	return nil
}

func (m *InMemoryWeatherRepository) GetByCity(city string) (*domain.WeatherData, error) {
	return m.collection[city], nil
}
