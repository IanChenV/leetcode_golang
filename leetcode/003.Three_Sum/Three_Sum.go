package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	start, end, index := 0, 0, 0
	result := make([][]int, 0)
	nums_size := len(nums)
	add_num := 0

	for index = 1; index < nums_size; index++ {
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for index > start && index < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if end < nums_size-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			add_num = nums[start] + nums[index] + nums[end]
			if add_num == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if add_num > 0 {
				end--
			} else if add_num < 0 {
				start++
			}
		}
		return result
	}

}
