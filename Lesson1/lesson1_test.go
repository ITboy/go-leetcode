package lesson1

import (
	"testing"
)

type testcase struct {
	nums   []int
	sum    int
	expect []int
}

func TestTwoSum(t *testing.T) {
	cases := []testcase{
		{nums: []int{1, 2, 3, 4, 5}, sum: 5, expect: []int{0, 4}},
		{nums: []int{}, sum: 1, expect: nil},
		{nums: []int{3, 6, 9}, sum: 4, expect: nil},
	}

	for _, c := range cases {
		got := twoSum(c.nums, c.sum)
		if got == nil {
			if c.expect != nil {
				t.Errorf("got twoSum(%v, %v) = %v, expect %v", c.nums, c.sum, got, c.expect)
			}
		} else if got != nil && c.nums[got[0]]+c.nums[got[1]] != c.sum {
			t.Errorf("got twoSum(%v, %v) = %v, but %v+%v != %v", c.nums, c.sum, got, c.nums[got[0]], c.nums[got[1]], c.sum)
		}
	}
}
