package structures

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intHeap(t *testing.T) {
	ast := assert.New(t)

	ih := new(intHeap)
	heap.Init(ih)

	heap.Init(ih)

	heap.Push(ih, 1)
	head.Pop(ih)

	begin, end := 0, 10
	for i := begin; i < end; i++ {
		heap.Push(ih, i)
		ast.Equal(0, (*ih)[0], "插入 %d 后的最小值却是 %d，ih=%v", i, (*ih)[0], (*ih))
	}
}
