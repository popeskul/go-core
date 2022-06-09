package eldest

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	args []Person
	want Person
}{
	{
		name: "Test 1: all types are Employee",
		args: []Person{
			Employee{
				age: 20,
			},
			Employee{
				age: 30,
			},
		},
		want: Employee{
			age: 30,
		},
	},
	{
		name: "Test 2: all types are Customer",
		args: []Person{
			Customer{
				age: 20,
			},
			Customer{
				age: 30,
			},
		},
		want: Customer{
			age: 30,
		},
	},
}

func TestEldest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.args {
				got := EldestWithSwitch(p)

				if !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
					t.Errorf("got = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
				}
			}
		})
	}
}

func TestEldestWithGenerics(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EldestWithGenerics(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEldestWithSwitch(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.args {
				got := EldestWithSwitch(p)

				if !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
					t.Errorf("got = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(tt.want))
				}
			}
		})
	}
}

// Faster method:
// #1 - if
// #2 - switch
// #3 - generics
func BenchmarkEldest(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Eldest(test.args)
			}
		})
	}
}

func BenchmarkEldestWithSwitch(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EldestWithSwitch(test.args)
			}
		})
	}
}

func BenchmarkEldestWithGenerics(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EldestWithGenerics(test.args...)
			}
		})
	}
}
