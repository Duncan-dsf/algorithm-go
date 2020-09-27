package normal

func sumOfLeftLeaves(root *TreeNode) int {

	res := 0
	var dfs func(root *TreeNode, isLeft bool)
	dfs = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if isLeft {
				res += root.Val
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, true)
		}
		if root.Right != nil {
			dfs(root.Right, false)
		}
	}

	dfs(root, false)
	return res
}
