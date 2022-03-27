package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	cost := 0
    for carsCount >= 0 {
		if carsCount >= 10 {
			cost += 95000
			carsCount -= 10
		} else if carsCount <= 9 {
			cost += carsCount * 10000
			break
		}
    }
	return uint(cost)
}
