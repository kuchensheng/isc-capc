package util

import "testing"

func TestRandString(t *testing.T) {
	type args struct {
		lenNum int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "stringtest",
			args: args{8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandString(tt.args.lenNum)
			t.Log(got)
		})
	}
}
