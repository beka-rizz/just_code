package hw_1

import "testing"

func Test_equal1(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Two arrays are equal",
			args{
				[]int{1, 2, 3},
				[]int{1, 2, 3},
			},
			true,
		},
		{"Two arrays have different sizes",
			args{
				[]int{1, 2, 3, 4},
				[]int{1, 2, 3},
			},
			false,
		},
		{"Two arrays are not equal",
			args{
				[]int{1, 2, 4},
				[]int{1, 2, 3},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equal1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equal2(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Two arrays are equal",
			args{
				[]int{5, 6, 9, 7},
				[]int{6, 7, 9, 5},
			},
			true,
		},
		{"Two arrays have different sizes",
			args{
				[]int{1, 2, 3, 4, 1},
				[]int{1, 2, 3, 4},
			},
			false,
		},
		{"Two arrays are not equal",
			args{
				[]int{2, 1, 3, 5, -1},
				[]int{1, 2, 3, -1, 2},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equal2() = %v, want %v", got, tt.want)
			}
		})
	}
}
