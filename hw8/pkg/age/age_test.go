package age

import (
	"testing"
)

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name string
		args []Person
		want int
	}{
		{
			name: "Test1: all types are Employee and Customer",
			args: []Person{
				Employee{
					age: 20,
				},
				Customer{
					age: 30,
				},
			},
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args...); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
