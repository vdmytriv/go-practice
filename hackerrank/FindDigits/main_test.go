package FindDigits

import "testing"

func Test_findDigits(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"simple test",
			args{1012},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDigits(tt.args.n); got != tt.want {
				t.Errorf("findDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
