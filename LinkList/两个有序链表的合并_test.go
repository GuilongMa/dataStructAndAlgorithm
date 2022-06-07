package LinkList

import "testing"

func TestLinkList_Merge(t *testing.T) {
	type args struct {
		l2 *LinkList
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_merge",
			args: args{l2: NewLinkList([]string{"1", "5"})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Merge(tt.args.l2); got != nil {
				t.Logf("Merge() = %v", got)
			}
		})
	}
}
