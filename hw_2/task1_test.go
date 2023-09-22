package hw_2

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{
				13,
			},
			"XIII",
		},
		{
			"Test 2",
			args{
				1994,
			},
			"MCMXCIV",
		},
		{
			"Test 3",
			args{
				3999,
			},
			"MMMCMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
