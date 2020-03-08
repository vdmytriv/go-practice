package CutTheSticks

import (
	"reflect"
	"testing"
)

func Test_cutTheSticks(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			"simple test",
			args{[]int32{5, 4, 4, 2, 2, 8}},
			[]int32{6, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutTheSticks(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cutTheSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}
