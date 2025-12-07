package lasagna

// OvenTime is the number of minutes the lasagna should be in the oven.
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes the lasagna needs to stay in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the preparation time based on the number of layers.
// Each layer takes 2 minutes to prepare.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time already spent: preparation + time in oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}