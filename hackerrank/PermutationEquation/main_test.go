package main

import (
	"reflect"
	"testing"
)

func Test_permutationEquation(t *testing.T) {
	type args struct {
		p []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			"simple test",
			args{[]int32{5, 2, 1, 3, 4}},
			[]int32{4, 2, 5, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutationEquation(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutationEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
