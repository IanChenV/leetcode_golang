package leetcode

import (
	"fmt"
	"leetcode_golang/structure"
	"testing"
)

type question002 struct {
	para2
	ans2
}

type para2 struct {
	one     []int
	another []int
}

type ans struct {
	one []int
}

func Test_Question2(t *testing.T) {
	qs := []question002{}
	fmt.Printf("Leetcode Problem 2\n")
	for _, q := range qs {
		_, p := q.ans2, q.para2
		fmt.Printf("【input】: %v	【output】: %v\n", p, structure.List2Ints(addTwoNumbers(structure.Ints2List(p.one), structure.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
