package lesson4

type solution3 struct{}

func (s solution3) findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := m + n
	if k%2 == 1 {
		return float64(findKthElement(nums1, nums2, k/2+1))
	} else {
		return float64(findKthElement(nums1, nums2, k/2)+findKthElement(nums1, nums2, k/2+1)) / 2
	}
}

func findKthElement(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	} else if len(nums2) == 0 {
		return nums1[k-1]
	} else if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	m := min(k/2-1, len(nums1)-1, len(nums2)-1)
	if nums1[m] < nums2[m] {
		nums1 = nums1[m+1:]
		return findKthElement(nums1, nums2, k-m-1)
	} else if nums1[m] > nums2[m] {
		nums2 = nums2[m+1:]
		return findKthElement(nums1, nums2, k-m-1)
	} else {
		if k%2 == 0 && m == (k/2-1) {
			return nums1[m]
		} else {
			nums1 = nums1[m+1:]
			nums2 = nums2[m+1:]
			return findKthElement(nums1, nums2, k-2*m-2)
		}
	}
}

func min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	} else {
		return min(nums[0], min(nums[1:]...))
	}
}
