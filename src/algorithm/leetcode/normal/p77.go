package normal

func combine(n int, k int) [][]int {

	var temp []int
	var res [][]int
	c := 1

	for i:=n; i>k; i-- {
		c *= i
	}
	dfs(&temp, 1, k, n, &res)
	return res
}

func dfs(temp *[]int, cur, remain, n int, res *[][]int) {

	if remain == 0 {
		cp := make([]int, len(*temp))
		copy(cp, *temp)
		*res = append(*res, cp)
		return
	}
	if cur > n {
		return
	}
	for i:=cur; i<=n; i++ {
		*temp = append(*temp, i)
		dfs(temp, i+1, remain-1, n, res)
		*temp = (*temp)[:len(*temp)-1]
	}
}

func P77(n int, k int) [][]int {
	return combine(n, k)
}