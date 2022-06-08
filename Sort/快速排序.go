package Sort

func QuickSort(array []int, l, r int) []int {
	if l >= r {
		return array
	}
	index := GetPivot(array, l, r)
	QuickSort(array, l, index-1)
	QuickSort(array, index+1, r)
	return array
}

func GetPivot(array []int, l, r int) (index int) {
	val := array[r]
	j := l
	for i := l; i < r; i++ {
		if array[i] < val {
			tmp := array[i]
			array[i] = array[j]
			array[j] = tmp
			j += 1
		}
	}
	array[r] = array[j]
	array[j] = val
	return j
}
