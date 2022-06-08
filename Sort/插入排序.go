package Sort

func InsertSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		tmp := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > tmp {
				array[j+1] = array[j]
				continue
			}
			break
		}
		array[j+1] = tmp
	}
	return array
}
