package weather

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type Weather interface {
	GetTemperature(location string) (Temperature, error)
}
