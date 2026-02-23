package exercise3

//Write a function called ArrStats accepting an []int and returning 3 values: min, max, average
func ArrStats(arr []int) (min int, max int, avg float64) {
	if len(arr) == 0 {
		return 0, 0, 0.
	}

	sum := arr[0]
	min, max = arr[0], arr[0]

	for _, val := range arr[1:] {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		sum += val
	}

	return min, max, float64(sum) / float64(len(arr))
}
