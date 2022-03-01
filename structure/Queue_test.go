package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.IsEmpty(), "检查新建的q是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		ast.Equal(i-start+1, q.Len(), "Push后检查q的长度")
	}

	for i := start; i < end; i++ {
		ast.Equal(i, q.Pop(), "从 q 中pop 出数来。")
	}

	ast.True(q.IsEmpty(), "检查 Pop完毕后的去是否为空")
}
