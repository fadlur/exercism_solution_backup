package cards

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	} else {
		return slice[index]
	}
}

func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

func PrependItems(slice []int, values ...int) []int {
	if values != nil {
		newData := []int{}
		for _, v := range values {
			newData = append(newData, v)
		}
		newData = append(newData, slice...)
		return newData
	} else {
		return slice
	}
}

func RemoveItem(slice []int, index int) []int {
	newData := []int{}
	for i, v := range slice {
		if i != index {
			newData = append(newData, v)
		}
	}
	return newData
}
