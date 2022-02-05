// Package weather provides weather information by location.
package weather

// CurrentCondition provides actual weather information.
var CurrentCondition string

// CurrentLocation provides weather information by place.
var CurrentLocation string

//Forecast function from the condition and city variables of the weather package //returns the display of the current position as well as the current meteorological //conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
