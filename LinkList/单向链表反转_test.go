package LinkList

import (
	"testing"
)

func TestLinkList_Reverse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_link_list_reverse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Reverse(); got != nil {
				t.Logf("Reverse() = %v", got)
			}
		})
	}
}
