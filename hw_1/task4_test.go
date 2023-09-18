package hw_1

import (
	"reflect"
	"testing"
)

func TestMySlice_sort(t *testing.T) {
	type fields struct {
		arr []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			"Test 1",
			fields{
				[]int{34, 43, 23, 76, 45, 12, 0, 2, 5},
			},
			[]int{0, 2, 5, 12, 23, 34, 43, 45, 76},
		},
		{
			"Test 2",
			fields{
				[]int{76, 45, -12, 0, 2, 5},
			},
			[]int{-12, 0, 2, 5, 45, 76},
		},
		{
			"Test 3",
			fields{
				[]int{77, 77, 3, 4, 0, 3, 1, 2},
			},
			[]int{0, 1, 2, 3, 3, 4, 77, 77},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MySlice{
				arr: tt.fields.arr,
			}
			if got := ms.sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
