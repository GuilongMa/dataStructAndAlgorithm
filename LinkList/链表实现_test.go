package LinkList

import (
	"testing"
)

var l *LinkList

func TestLinkList_BackInsert(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_back_insert",
			args: args{data: "1"},
		},
		{
			name: "test_link_list_back_insert",
			args: args{data: "2"},
		},
		{
			name: "test_link_list_back_insert",
			args: args{data: "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.BackInsert(tt.args.data)
		})
	}
}

func TestLinkList_Delete(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_delete",
			args: args{data: "2"},
		},
		{
			name: "test_link_list_delete",
			args: args{data: "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Delete(tt.args.data)
		})
	}
}

func TestLinkList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "test_link_list_is_empty",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_PreInsert(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_pre_insert",
			args: args{
				data: "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.PreInsert(tt.args.data)
		})
	}
}

func TestLinkList_PrintAllElements(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_link_list_print_all_elements",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf(l.String())
		})
	}
}

func TestLinkList_Search(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_link_list_search",
			args: args{data: "2"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l.Search(tt.args.data); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkList_Update(t *testing.T) {

	type args struct {
		oldData interface{}
		newData interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_update",
			args: args{
				oldData: "3",
				newData: "4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Update(tt.args.oldData, tt.args.newData)
		})
	}
}

func TestNewLinkList(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "test_link_list",
			args: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if l = NewLinkList(tt.args); l != nil {
				t.Logf("NewLinkList() = %v, ", l)
			}
		})
	}
}

func TestNode_String(t *testing.T) {
	type fields struct {
		data interface{}
		next *Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test_string",
			fields: fields{
				data: "1",
				next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				data: tt.fields.data,
				next: tt.fields.next,
			}
			node.String()
		})
	}
}

func TestAll(t *testing.T) {

	t.Run("New LinkList", TestNewLinkList)
	// insert 0
	t.Run("Pre Insert", TestLinkList_PreInsert)
	// insert 1,2,3
	t.Run("Back Insert", TestLinkList_BackInsert)
	// update 3 to 4
	t.Run("Update", TestLinkList_Update)
	// search 2
	t.Run("Search", TestLinkList_Search)
	// delete 1,2
	t.Run("Delete", TestLinkList_Delete)
	// is empty
	t.Run("Is Empty", TestLinkList_IsEmpty)
	// print all elements
	t.Run("Print All Elements", TestLinkList_PrintAllElements)
	// merge
	t.Run("Merge", TestLinkList_Merge)
	// get middle point
	t.Run("Get Mid Point", TestLinkList_MidPoint)
	// has cycled
	t.Run("Has Cycle", TestLinkList_HasCycle)
	// delete last n node
	t.Run("Delete Last N Node", TestLinkList_DeleteNode)
	// reverse
	t.Run("Reverse", TestLinkList_Reverse)
	// palindrome
	t.Run("Palindrome", TestLinkListPalindrome)

}
