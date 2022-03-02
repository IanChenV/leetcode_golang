package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, found := m[v]; found {
			return true
		}
		m[v] = true
	}
	return false
}
