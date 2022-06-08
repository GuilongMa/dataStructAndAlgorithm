package Sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_insert_sort",
			args: args{array: []int{4, 5, 6, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("src = %v", tt.args.array)
			got := InsertSort(tt.args.array)
			t.Logf("InsertSort() = %v, want %v", got, tt.want)
		})
	}
}
