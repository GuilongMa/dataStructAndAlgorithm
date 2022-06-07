package LinkList

import "testing"

func TestLinkList_DeleteNode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_delete_n_node",
			args: args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.DeleteNode(tt.args.n)
			t.Logf("after delete:%v", l)
		})
	}
}
