package normal

func combinationSum3(k int, n int) [][]int {

	var dfs func(i, c, sum int, temp []int, res *[][]int)
	dfs = func(i, c, sum int, temp []int, res *[][]int) {
		if sum == n && c == k {
			t := make([]int, len(temp))
			copy(t, temp)
			*res = append(*res, t)
			return
		}
		if i == 10 || sum > n || c > k {
			return
		}
		for j := i; j < 10; j++ {
			if sum+j > n {
				break
			}
			temp = append(temp, j)
			dfs(j+1, c+1, sum+j, temp, res)
			temp = temp[:len(temp)-1]
		}
	}

	var temp []int
	var res [][]int
	dfs(1, 0, 0, temp, &res)
	return res
}
