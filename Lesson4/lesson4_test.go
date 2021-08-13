package lesson4

import (
	"math"
	"testing"
)

type testcase struct {
	nums1, nums2 []int
	expected     float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []testcase{
		{nums1: []int{1}, nums2: []int{3}, expected: 2},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, expected: 2.5},
		{nums1: []int{0, 0}, nums2: []int{0, 0}, expected: 0},
		{nums1: []int{}, nums2: []int{1}, expected: 1},
		{nums1: []int{2}, nums2: []int{}, expected: 2},
		{nums1: []int{1}, nums2: []int{2, 3, 4, 5, 6}, expected: 3.5},
		{nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}, nums2: []int{0, 6}, expected: 10.5},
	}
	solutions := []solution{solution1{}, solution2{}, solution3{}}
	for _, s := range solutions {
		for i, c := range cases {
			if got := s.findMedianSortedArrays(c.nums1, c.nums2); math.Abs(got-c.expected) > 0.00001 {
				t.Errorf("%T.findMedianSortedArrays(%v, %v)=%v, expected: %v", s, c.nums1, c.nums2, got, c.expected)
			} else {
				t.Logf("%T case %v passed, %v\n", s, i, c)
			}
		}
	}
}
