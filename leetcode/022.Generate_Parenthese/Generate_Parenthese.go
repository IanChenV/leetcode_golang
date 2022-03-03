package leetcode


func 

func dfs(lindex int, rindex int, str string res *[]string){
	if lindex == 0 && rindex ==0{
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		dfs(lindex-1, rindex, str + "(", res)
	}

	if rindex > 0{
		dfs(lindex, rindex -1, str + ")", res)
	}
}