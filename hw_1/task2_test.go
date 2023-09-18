package hw_1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1",
			args{
				[]string{"aaa", "aa", "aaa"},
			},
			"aa",
		},
		{"Test 2",
			args{
				[]string{"flower", "flow", "flutter"},
			},
			"fl",
		},
		{"Test 3",
			args{
				[]string{"ananas", "anas", "animal"},
			},
			"an",
		},
		{"Test 4",
			args{
				[]string{"z"},
			},
			"z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
