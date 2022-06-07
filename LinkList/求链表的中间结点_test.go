package LinkList

import (
	"testing"
)

func TestLinkList_MidPoint(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_link_list_mid_point",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSlow := l.MidPoint(); gotSlow != nil {
				t.Logf("MidPoint() = %v", gotSlow)
			}
		})
	}
}
