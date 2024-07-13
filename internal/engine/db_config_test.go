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
		name      string
		args      args
		want      DBConfig
		shouldErr bool
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
			shouldErr: false,
		},
		{
			name: "should err for incorrect connection string",
			args: args{
				url: "postgres://user:pass@host:port",
			},
			want:      DBConfig{},
			shouldErr: true,
		},
		{
			name: "should err for empty connection string",
			args: args{
				url: "",
			},
			want:      DBConfig{},
			shouldErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := configFromURL(tt.args.url)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("configFromURL() should have returned err")
				}
			} else {
				if !reflect.DeepEqual(tt.want, got) {
					t.Errorf("configFromURL() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
