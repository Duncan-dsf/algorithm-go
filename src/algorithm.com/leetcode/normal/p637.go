package normal

import "container/list"

func averageOfLevels(root *TreeNode) []float64 {

	l1, l2 := list.New(), list.New()
	l1.PushFront(root)
	res := make([]float64, 0)
	for l1.Len() > 0 {
		sum, c := 0, 0
		for node := l1.Front(); node != nil; node = node.Next() {

			treeNode := (*node).Value.(*TreeNode)
			c++
			sum += treeNode.Val
			if treeNode.Left != nil {
				l2.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				l2.PushBack(treeNode.Right)
			}
		}
		l1, l2 = l2, l1
		l2.Init()
		res = append(res, float64(sum)/float64(c))
	}
	return res
}
