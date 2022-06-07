package LinkList

import "testing"

func TestLinkListPalindrome(t *testing.T) {
	type args struct {
		l *LinkList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_link_list_palindrome_1",
			args: args{l: NewLinkList([]string{"1", "2", "4", "2", "1"})},
			want: true,
		},
		{
			name: "test_link_list_palindrome_2",
			args: args{l: NewLinkList([]string{"1", "2", "2", "1"})},
			want: true,
		},
		{
			name: "test_link_list_palindrome_3",
			args: args{l: NewLinkList([]string{"1", "2"})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Palindrome(tt.args.l)
			t.Logf("Palindrome() = %v, want %v", got, tt.want)
		})
	}
}
