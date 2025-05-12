package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{nums: []int{1, 5, 3, 7}, target: 8, want: []int{0, 3}},
		{nums: []int{-1, -2, -3, -4, -5}, target: -8, want: []int{2, 4}},
		{nums: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
		{nums: []int{1, 2}, target: 3, want: []int{0, 1}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		// Since the answer can be in any order, check both possibilities
		if !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, []int{tt.want[1], tt.want[0]}) {
			t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
