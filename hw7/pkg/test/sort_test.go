package test

import (
	"reflect"
	"sort"
	"testing"
)

func TestSingleSortString(t *testing.T) {
	got := []string{"123", "*(&^*(", "-", "", "some", "_______"}
	want := []string{"", "*(&^*(", "-", "123", "_______", "some"}

	sort.Strings(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v, want %v", got, want)
	}
}

func TestSortString(t *testing.T) {
	type args struct {
		s []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test with strings",
			args: args{
				s: []string{"123", "*(&^*(", "-", "", "some", "_______"},
			},
			want: []string{"", "*(&^*(", "-", "123", "_______", "some"},
		},
		{
			name: "Test with empty slice",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.args.s)
			got := tt.args.s

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortInt(t *testing.T) {
	type args struct {
		n []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test with ints",
			args: args{
				n: []int{20, -50, 0, 100, -2, 500},
			},
			want: []int{-50, -2, 0, 20, 100, 500},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.args.n)
			got := tt.args.n

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
