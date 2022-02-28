package structure

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (st *Stack) Pop() int {
	res := st.nums[len(st.nums)-1]
	st.nums = st.nums[:len(st.nums)-1]
	return res
}

func (st *Stack) Push(n int) {
	st.nums = append(st.nums, n)
}

func (st *Stack) Len() int {
	return len(st.nums)
}

func (st *Stack) IsEmpty() bool {
	return st.Len() == 0
}
