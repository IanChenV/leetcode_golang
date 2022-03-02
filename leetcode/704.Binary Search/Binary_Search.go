package leetcode

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1

}
