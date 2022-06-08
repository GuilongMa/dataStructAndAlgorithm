package StackAndQueue

import (
	"fmt"
	"testing"
)

var q = NewLinkListQueue()

func TestLinkListQueue_Get(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_link_list_queue_get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := q.Get(); got != nil {
				t.Logf("Get() = %v", got)
				fmt.Println(q)
			}
		})
	}
}

func TestLinkListQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "test_link_list_queue_is_empty",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q.IsEmpty()
			t.Logf("IsEmpty() = %v, want %v", got, tt.want)
		})
	}
}

func TestLinkListQueue_Push(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_link_list_queue_push_1",
			args: args{data: "1"},
		},
		{
			name: "test_link_list_queue_push_2",
			args: args{data: "2"},
		},
		{
			name: "test_link_list_queue_push_3",
			args: args{data: "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q.Push(tt.args.data)
			fmt.Println(q)
		})
	}
}

func TestImplement(t *testing.T) {
	t.Run("Push", TestLinkListQueue_Push)
	t.Run("IsEmpty", TestLinkListQueue_IsEmpty)
	t.Run("Get", TestLinkListQueue_Get)
}
