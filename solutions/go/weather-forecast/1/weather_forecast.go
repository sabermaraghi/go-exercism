//Package weather provides tools to get the current state of the weather and the city and show them in a beautiful format.
package weather

var (
    // CurrentCondition represents a current state of the weather.
	CurrentCondition string
    // CurrentLocation represents the current location where the weather should be calculated.
	CurrentLocation  string
)
// Forecast function returns the current weather of the current location, relying on the input city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
