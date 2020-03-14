package main

import "testing"

func Test_factorial(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple test: 25",
			args{25},
			"15511210043330985984000000",
		},
		{
			"simple test: 45",
			args{45},
			"119622220865480194561963161495657715064383733760000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.n); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
