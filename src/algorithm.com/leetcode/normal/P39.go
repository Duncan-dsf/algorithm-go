package normal

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var dfs func(i, s int, temp []int, res *[][]int)
	dfs = func (i, s int, temp []int, res *[][]int) {

		if s == target {
			t := make([]int, len(temp))
			copy(t, temp)
			*res = append(*res, t)
			return
		}
		if s > target {
			return
		}
		for j:=i; j<len(candidates); j++ {
			v := candidates[j]
			if s+v > target {
				break
			}
			temp = append(temp, v)
			dfs(j, v+s, temp, res)
			temp = temp[:len(temp)-1]
		}

	}
	var res [][]int
	var temp []int
	dfs(0, 0, temp, &res)
	return res
}

func P39(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)
}
