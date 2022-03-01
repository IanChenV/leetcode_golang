package structure

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_priorityQueue(t *testing.T) {
	ast := assert.New(t)

	items := map[string]int{
		"banana": 2,
		"apple":  1,
		"pear":   3,
	}

	pq := make(PQ, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &entry{
			key:      value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	it := &entry{
		key:      "orange",
		priority: 5,
	}
	heap.Push(&pq, it)
	pq.update(it, it.key, 0)

	expected := []string{
		"orange",
		"apple",
		"banana",
		"pear",
	}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*entry)
		ast.Equal(expected[it.priority], it.key)
	}
}
