package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var main [][]int
	if root == nil {
		return main
	}
	var temp = []*Node{root}
	//main = append(main, []int{root.Val})
	for len(temp) > 0 {
		var c []int
		var nt []*Node
		for _, i := range temp {
			c = append(c, i.Val)
			for _, j := range i.Children {
				nt = append(nt, j)
			}
		}
		main = append(main, c)
		temp = nt
	}
	return main
}
