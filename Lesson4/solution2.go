package lesson4

type solution2 struct{}

func (s solution2) findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	len3 := len1 + len2
	if len1 == 0 {
		if len2%2 == 1 {
			return float64(nums2[len2/2])
		} else {
			return float64((nums2[len2/2-1] + nums2[len2/2]) / 2)
		}
	} else if len2 == 0 {
		if len1%2 == 1 {
			return float64(nums1[len1/2])
		} else {
			return float64((nums1[len1/2-1] + nums1[len1/2]) / 2)
		}
	} else {
		i, j := 0, 0
		m, n := 0, 0
		for k := 0; k <= len3/2; {
			if j == len2 || (i < len1 && j < len2 && nums1[i] < nums2[j]) {
				m = n
				n = nums1[i]
				i++
				k++

			} else if i == len1 || (i < len1 && j < len2 && nums1[i] >= nums2[j]) {
				m = n
				n = nums2[j]
				j++
				k++
			}
		}
		if len3%2 == 1 {
			return float64(n)
		} else {
			return float64(m+n) / 2
		}
	}
}
