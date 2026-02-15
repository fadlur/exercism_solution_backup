package lasagna

// TODO: define the 'PreparationTime()' function
// layers := []string{"sauce", "noddles", "sauce", "meat", "mozzarella", "noodles"}
func PreparationTime(layers []string, averageTime int) int {
    if (averageTime == 0) {
        return len(layers) * 2
    }
	return len(layers) * averageTime
    
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (totalWeight int, totalSauce float64) {
    totalWeight = 0
    totalSauce = 0.0
    for _, layer := range(layers) {
        if (layer == "noodles") {
            totalWeight += 50
        } else if (layer == "sauce") {
            totalSauce += 0.2
        }
    }    
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
    newQuantities := make([]float64, len(quantities))
    for i, qty := range(quantities) {
        newQuantities[i] = qty * (float64(portion)/2)
    }
    return newQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
