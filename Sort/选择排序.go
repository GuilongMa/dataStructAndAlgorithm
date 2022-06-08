package Sort

func ChooseSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		tmp := array[i]
		array[i] = array[min]
		array[min] = tmp
	}
	return array
}
