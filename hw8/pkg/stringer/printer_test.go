package stringer

import (
	"bytes"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{
			name: "Test 1: strings",
			args: []interface{}{"string", " ", "string2"},
			want: "string string2",
		},
		{
			name: "Test 2: strings and others",
			args: []interface{}{"string", 85, nil, " ", "string2"},
			want: "string string2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := Write(&buf, tt.args...)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if got := buf.String(); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
