package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {

	end1, end2 := m-1, n-1
	end := m + n - 1
	for end1 >= 0 && end2 >= 0 && end >= 0 {
		if nums1[end1] >= nums2[end2] {
			nums1[end] = nums1[end1]
			end1--
		} else {
			nums1[end] = nums2[end2]
			end2--
		}
		end--
	}

	for end2 >= 0 {
		nums1[end] = nums2[end2]
		end2--
		end--
	}
}
