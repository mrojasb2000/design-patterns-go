package facade

type CurrentWeatherDataRetriever interface {
	// Acceptance criteria 2
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	// Acceptance criteria 3
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}
