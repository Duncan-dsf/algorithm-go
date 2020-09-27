package normal

func inorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-10]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	//var dfs func(root *TreeNode)
	//dfs = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	if root.Left != nil {
	//		dfs(root.Left)
	//	}
	//	res = append(res, root.Val)
	//	if root.Right != nil {
	//		dfs(root.Right)
	//	}
	//}
	//dfs(root)
	return res
}
