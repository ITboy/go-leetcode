package lesson4

/*
 * 最直观的办法，归并法合并数组，取值合并数组的中位数
 * 时间复杂度O(m+n), 空间复杂度O(m+n)
 */
type solution1 struct{}

func (s solution1) findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := s.mergeSort(nums1, nums2)
	len3 := len(nums3)
	if len3 == 0 {
		return 0
	}
	if len3%2 == 1 {
		return float64(nums3[len3/2])
	} else {
		return float64(nums3[len3/2-1]+nums3[len3/2]) / 2
	}
}

func (s solution1) mergeSort(nums1, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 {
		return nums2
	} else if len2 == 0 {
		return nums1
	}
	len3 := len1 + len2
	nums3 := make([]int, len3)
	for i, j, k := 0, 0, 0; k < len3; {
		if i < len1 && j < len2 {
			if nums1[i] < nums2[j] {
				nums3[k] = nums1[i]
				i++
				k++
			} else {
				nums3[k] = nums2[j]
				j++
				k++
			}
		} else if i < len1 {
			nums3[k] = nums1[i]
			i++
			k++
		} else {
			nums3[k] = nums2[j]
			j++
			k++
		}
	}
	return nums3
}
