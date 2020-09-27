package _2020fall

/*
TODO 没过
*/
func minimumOperations(leaves string) int {

	n, y := len(leaves), 0
	for i:=0; i<n; i++ {
		if leaves[i] == 'y' {
			y++
		}
	}

	res := y
	for l:=1; l+y < n; l++ {
		c := 0
		for i:=0; i<y; i++ {
			if leaves[l+i] != 'y' {
				c++
			}
		}
		if res > c {
			res = c
		}
	}
	return res
}

func P3(leaves string) int {
	return minimumOperations(leaves)
}