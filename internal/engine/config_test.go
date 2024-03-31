package engine

import (
	"reflect"
	"testing"
)

func Test_configFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name        string
		args        args
		want        DBConfig
		shouldPanic bool
	}{
		{
			name: "should parse valid connection string",
			args: args{
				url: "postgres://user:pass@host:port/dbname",
			},
			want: DBConfig{
				User:     "user",
				Password: "pass",
				Host:     "host",
				Port:     "port",
				Name:     "dbname",
			},
			shouldPanic: false,
		},
		{
			name: "should panic for incorrect connection string",
			args: args{
				url: "postgres://user:pass@host:port",
			},
			want:        DBConfig{},
			shouldPanic: true,
		},
		{
			name: "should panic for empty connection string",
			args: args{
				url: "",
			},
			want:        DBConfig{},
			shouldPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			if got := configFromURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
