package repository

import (
	"reflect"
	"testing"
)

func TestNewPostgresDB(t *testing.T) {
	tests := []struct {
		name  string
		args  Config
		want  *DB
		error bool
	}{
		{
			name: "Success",
			args: Config{
				User:     "postgres",
				Password: "postgres",
				Url:      "localhost",
				Port:     "5432",
				DBName:   "postgres",
			},
			want:  &DB{},
			error: false,
		},
		{
			name: "Fail",
			args: Config{
				User:     "asd",
				Password: "asd",
				Url:      "locafsalhost",
				Port:     "234",
				DBName:   "asd",
			},
			want:  &DB{},
			error: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresDB(tt.args)

			if (err != nil) != tt.error {
				t.Errorf("NewPostgresDB() error = %v, wantErr %v", err, tt.error)
			}

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewPostgresDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
