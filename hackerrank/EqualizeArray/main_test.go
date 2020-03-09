package EqualizeArray

import "testing"

func Test_equalizeArray(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"simple test",
			args{[]int32{3, 3, 2, 1, 3}},
			int32(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalizeArray(tt.args.arr); got != tt.want {
				t.Errorf("equalizeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
