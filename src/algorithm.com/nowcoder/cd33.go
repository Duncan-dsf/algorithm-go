package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.nowcoder.com/practice/210741385d37490c97446aa50874e62d
// TODO 95
func lcs(string1, string2 string) (string, int) {

	s1, s2 := []rune(string1), []rune(string2)
	n, m := len(s1), len(s2)
	if n == 0 || m == 0 {
		return "-1", 0
	}
	dp := make([][]int, n)
	var res []rune
	var max int
	for i:=0; i<n; i++ {
		dp[i] = make([]int, m)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {

			c := 0
			if s1[i] == s2[j] {
				c++
				if i>1 && j>1 && s1[i-1] == s2[j-1] {
					c += dp[i-1][j-1]
				}
			}
			dp[i][j] = c
			if max < c {
				max = c
				res = s1[i+1-c:i+1]
			}
		}
	}
	return string(res), max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		s, _ := lcs(s1, s2)
		if s == "" {
			fmt.Println("-1")
		} else {
			fmt.Println(s)
		}
	}
}