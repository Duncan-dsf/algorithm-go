package c206

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {

	m := map[int]int{}
	for _, pair := range pairs {
		m[pair[0]] = pair[1]
		m[pair[1]] = pair[0]
	}

	var check func(i int) bool
	check = func(x int) bool {

		y := m[x]
		for _, u := range preferences[x] {

			if u == y {
				return false
			}
			v := m[u]
			for _, t := range preferences[u] {
				if t == x {
					return true
				}
				if t == v {
					break
				}
			}
		}
		return false
	}
	res := 0
	for _, pair := range pairs {

		if check(pair[0]) {
			res++
		}
		if check(pair[1]) {
			res++
		}
	}
	return res
}

func P2(n int, preferences [][]int, pairs [][]int) int {
	return unhappyFriends(n, preferences, pairs)
}
