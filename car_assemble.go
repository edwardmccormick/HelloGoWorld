package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate /100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	return int(CalculateWorkingCarsPerHour(productionRate, successRate)/float64(60.0))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsBulk :=uint ((uint(carsCount)/uint(10)) * uint(95000))
	carsSingular :=uint ((uint(carsCount) - (uint(carsCount)/uint(10))*uint(10)) *uint(10000))
	return carsBulk + carsSingular
}

