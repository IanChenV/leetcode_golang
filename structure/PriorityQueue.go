package structure

import "container/heap"

type entry struct {
	key      string
	priority int
	index    int
}

type PQ []*entry

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往pq中放entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop从pq中取出
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// 更新优先级别
func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
