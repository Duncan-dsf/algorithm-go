package c206

func numSpecial(mat [][]int) int {

	n, m := len(mat), len(mat[0])
	flag1 := make([]bool, n)
	flag2 := make([]bool, m)
	for i:=0; i<n; i++ {
		c := 0
		for j:=0; j<m; j++ {
			if mat[i][j] == 1 {
				c++
			}
		}
		flag1[i] = c < 2
	}
	for j:=0; j<m; j++ {
		c := 0
		for i:=0; i<n; i++ {
			if mat[i][j] == 1 {
				c++
			}
		}
		flag2[j] = c < 2
	}
	res := 0
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if mat[i][j] == 1 {
				if flag1[i] && flag2[j] {
					res++
				}
				break
			}
		}
	}
	return res
}
