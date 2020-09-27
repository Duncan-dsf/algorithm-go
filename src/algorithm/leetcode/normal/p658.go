package normal

func findRedundantDirectedConnection(edges [][]int) []int {

	n := len(edges)
	father := make([]int, n+1)

	var union func(a, b int)
	var find func(a int) int
	union = func(a, b int) {
		if find(a) != find(b) {
			father[father[b]] = father[a]
		}
	}
	find = func(a int) int {
		if a != father[a] {
			father[a] = find(father[a])
		}
		return father[a]
	}

	for i:=n-1; i>=0; i-- {
		for i:=1; i<=n; i++ {
			father[i] = i
		}
		f := true
		for j:=0; j<n; j++ {
			if i == j {
				continue
			}
			if find(edges[j][1]) != edges[j][1] || edges[j][1] == find(edges[j][0]) {
				f = false
				break
			} else {
				union(edges[j][0], edges[j][1])
			}
		}
		if f {
			return edges[i]
		}
	}
	return nil
}