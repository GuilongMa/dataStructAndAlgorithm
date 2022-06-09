package BinarySearch

func BinarySearch(array []int, val int) (int, bool) {
	l := 0
	r := len(array) - 1
	for {
		if l > r {
			break
		}
		mid := (l + r) / 2
		if array[mid] == val {
			return mid, true
		}
		if array[mid] < val {
			l = mid + 1
			continue
		}
		r = mid - 1
	}
	return -1, false
}
