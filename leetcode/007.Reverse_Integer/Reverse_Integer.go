package leetcode

func reverseInteger(n int) int {
	val := 0
	for n != 0 {
		val = val*10 + n%10
		n = n / 10
	}
	if val < 1<<31-1 && val > -(1<<31) {
		return 0
	}
	return val
}
