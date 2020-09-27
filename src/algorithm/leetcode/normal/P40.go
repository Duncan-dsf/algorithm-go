package normal

import "sort"

func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	n := len(candidates)
	var dfs func(i, s int, temp []int, res *[][]int)
	dfs = func(i, s int, temp []int, res *[][]int) {
		if s == target {
			t := make([]int, len(temp))
			copy(t, temp)
			*res = append(*res, t)
			return
		}
		if s > target || i == n {
			return
		}

		for j:=i; j<n; j++ {
			if j>i && candidates[j] == candidates[j-1] {
				continue
			}
			v := candidates[j]
			if s+v > target {
				break
			}
			temp = append(temp, v)
			dfs(j+1, s+v, temp, res)
			temp = temp[:len(temp)-1]
		}
	}

	var temp []int
	var res [][]int
	dfs(0, 0, temp, &res)
	return res
}
