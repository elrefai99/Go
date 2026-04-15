package sort

// Bubble sort
func BubbleSort(y []int) []int {
	for i, _ := range y {
		for j, _ := range y {
			if y[i] > y[j] {
				y[j], y[i] = y[i], y[j]
			}
		}
	}
	return y
}

func BubbleSort_2(y []int) []int {
	len := len(y)

	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if y[j] > y[j+1] {
				// swap
				y[j], y[j+1] = y[j+1], y[j]
			}
		}
	}
	return y
}
