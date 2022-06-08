package Sort

import "testing"

func TestImplement(t *testing.T) {
	t.Run("bubble sort", TestBubbleSort)
	t.Run("choose sort", TestChooseSort)
	t.Run("insert sort", TestInsertSort)
	t.Run("quick sort", TestQuickSort)
}
