// Package weather provides tools to forecast the current conditions given a location.
package weather

// CurrentCondition is a string describing the current conditions.
var CurrentCondition string
// CurrentLocation is a string describing the location.
var CurrentLocation string

// Forecast takes a city and condition string and returns a string forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
