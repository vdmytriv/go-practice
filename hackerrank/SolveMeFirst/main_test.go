package main

import "testing"

func Test_solveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"simple test",
			args{
				uint32(2),
				uint32(3),
			},
			uint32(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
