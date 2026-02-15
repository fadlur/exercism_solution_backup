package purchase

import "fmt"

func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option2)
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
		return originalPrice * 80 / 100
	}
	if age >= 10 {
		return originalPrice * 50 / 100
	}
	return originalPrice * 70 / 100
}
