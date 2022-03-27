// Package weather is a package about weather.
package weather

// CurrentCondition is CurrentCondition.
var CurrentCondition string

// CurrentLocation is CurrentLocation.
var CurrentLocation string

// Forecast is a function that gives current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
