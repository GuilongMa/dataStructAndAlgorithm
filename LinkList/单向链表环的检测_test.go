package LinkList

import "testing"

func TestLinkList_HasCycle(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "test_link_list_has_cycle",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := l.HasCycle()
			t.Logf("HasCycle() = %v, want %v", got, tt.want)
		})
	}
}
