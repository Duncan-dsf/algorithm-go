package normal

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p.Val > q.Val {
		p, q = q, p
	}
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {

		if r == nil {
			return nil
		}
		if r.Val >= p.Val && r.Val <= q.Val {
			return r
		}
		if r.Val < p.Val {
			return dfs(r.Right)
		} else {
			return dfs(r.Left)
		}
	}

	return dfs(root)
}

func P235(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}