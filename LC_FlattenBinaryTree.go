package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree(root *TreeNode) (start *TreeNode, end *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	var ls, lr, rs, rr *TreeNode
	var left, right bool
	if root.Left != nil {
		left = true
		ls, lr = tree(root.Left)
	}
	if root.Right != nil {
		right = true
		rs, rr = tree(root.Right)
	}
	if left && right {
		root.Left = nil
		root.Right = ls
		lr.Right = rs
		start = root
		end = rr
		return
	} else if left {
		root.Left = nil
		root.Right = ls
		start = root
		end = lr
		return
	} else if right {
		root.Left = nil
		root.Right = rs
		start = root
		end = rr
		return
	}
	return
}
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	_, _ = tree(root)
	return
}
