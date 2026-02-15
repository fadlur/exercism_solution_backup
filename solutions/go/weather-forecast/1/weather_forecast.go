// Package weather provides tools to forecast weather.
package weather

// CurrentCondition type is string.
var CurrentCondition string

// CurrentLocation type is string.
var CurrentLocation string

// Forecast takes 2 argument city and condition. both of them and the result are string type.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
