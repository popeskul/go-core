package geom

import (
	"testing"
)

func TestDistance(t *testing.T) {
	type params struct {
		x1, y1, x2, y2 float64
	}

	tests := []struct {
		name   string
		params params
		want   float64
	}{
		{
			name:   "Test 1",
			params: params{x1: 1, y1: 2, x2: 3, y2: 4},
			want:   2.8284271247461903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.params.x1, tt.params.y1, tt.params.x2, tt.params.y2)

			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
