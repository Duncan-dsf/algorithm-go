package normal

import "container/list"

type TreeNode struct {

	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	var res [][]int
	l1, l2 := list.New(), list.New()
	l1.PushBack(root)
	for l1.Len() > 0 {
		l2.Init()
		var temp []int
		for e := l1.Front(); e != nil; e = e.Next() {
			node := e.Value.(*TreeNode)
			temp = append(temp, (*node).Val)
			if (*node).Left != nil {
				l2.PushBack((*node).Left)
			}
			if (*node).Right != nil {
				l2.PushBack((*node).Right)
			}
		}
		l1, l2 = l2, l1
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
	}
	for l, r := 0, len(res)-1; l<r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
