package BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		array []int
		val   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "test_binary_search",
			args: args{
				array: []int{4, 5, 6, 3, 2, 1},
				val:   6,
			},
			want:  2,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearch(tt.args.array, tt.args.val)
			if got != tt.want {
				t.Errorf("BinarySearch() got = %v, want %v", got, tt.want)
			}
			t.Logf("BinarySearch() got = %v, want %v", got, tt.want)
			if got1 != tt.want1 {
				t.Errorf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
			t.Logf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
		})
	}
}
