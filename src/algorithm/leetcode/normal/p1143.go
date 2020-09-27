package normal

func longestCommonSubsequence(text1 string, text2 string) int {

	s1, s2 := []rune(text1), []rune(text2)
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i:=0; i<=n; i++ {
		dp[i] = make([]int, m+1)
	}
	//res := 0
	for i:=1; i<=n; i++ {
		for j:=1; j<=m; j++ {
			c := 0
			if s1[i-1] == s2[j-1] {
				c = dp[i-1][j-1] + 1
			} else {
				c = max(dp[i-1][j], dp[i][j-1])
			}
			dp[i][j] = c
			//res = max(c, res)
		}
	}
	return dp[n+1][m+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
