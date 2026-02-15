package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRatePerMinute := float64(productionRate) / 60
	return int(productionRatePerMinute * successRate / 100)
}

func CalculateCost(carsCount int) uint {
	return uint(carsCount/10)*95000 + (uint(carsCount)%10)*10000
}
