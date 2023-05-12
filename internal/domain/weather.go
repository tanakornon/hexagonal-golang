package domain

type WeatherData struct {
	City        string
	Temperature string
	Humidity    string
	Condition   string
}

type WeatherRepository interface {
	Save(weatherData *WeatherData) error
	GetByCity(city string) (*WeatherData, error)
}
