package main

import "testing"

func Test_viralAdvertising(t *testing.T) {
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
			args{5},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viralAdvertising(tt.args.n); got != tt.want {
				t.Errorf("viralAdvertising() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkViralAdvertising(b *testing.B) {
	for i := 0; i < b.N; i++ {
		viralAdvertising(100)
	}
}
