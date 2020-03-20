package main

import (
	"reflect"
	"testing"
)

func Test_serviceLane(t *testing.T) {
	type args struct {
		n     int32
		width []int32
		cases [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			"simple test",
			args{
				int32(8),
				[]int32{2, 3, 1, 2, 3, 2, 3, 3},
				[][]int32{
					{0, 3},
					{4, 6},
					{6, 7},
					{3, 5},
					{0, 7},
				},
			},
			[]int32{1, 2, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serviceLane(tt.args.n, tt.args.width, tt.args.cases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceLane() = %v, want %v", got, tt.want)
			}
		})
	}
}
