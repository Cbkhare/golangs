package main

import "fmt"

// using ds currently as array of int, later can modify it to other struct
type ds []int

type union ds

// qui is QuickUnionImproved
type qui struct {
	u union
	s ds
}

// QuickFind is an Interface
type QuickFind interface {
	QuickFindUnion(a int, b int)
	QuickFindConnected(a int, b int) bool
	length() int
	construct(n int)
}

// QuickUnion is an Interface
type QuickUnion interface {
	QuickUnionUnion(a int, b int)
	QuickUnionConnected(a int, b int) bool
	root(a int) int
	length() int
	construct(n int)
}

// QUImproved implement qui and is weighted implementation of Quick Union
type QUImproved interface {
	QuiUnion(a int, b int)
	QuiConnected(a int, b int) bool
	quiRoot(a int) int
	quiConstruct(n int)
}

// QUImprovedPathCompression implement qui and is weighted implementation of Quick Union
type QUImprovedPathCompression interface {
	QuiUnionPathComp(a int, b int)
	QuiConnectedPathComp(a int, b int) bool
	quiRootPathComp(a int) int
	quiConstruct(n int)
}

func (q *qui) quiRoot(a int) int {
	for (*q).u[a] != a {
		a = (*q).u[a]
	}
	return a // (*q).u[a]
}

func (q *qui) quiConstruct(n int) {
	for i := 0; i < n; i++ {
		(*q).u = append((*q).u, i)
		(*q).s = append((*q).s, 1)
	}
}

func (q *qui) QuiConnected(a int, b int) bool {
	return q.quiRoot(a) == q.quiRoot(b)
}

func (q *qui) QuiUnion(a int, b int) {
	a = q.quiRoot(a)
	b = q.quiRoot(b)
	// Root with shorter DFS joins the Larger one
	if (*q).s[a] > (*q).s[b] {
		(*q).u[b] = a
		(*q).s[a] += (*q).s[b]
	} else {
		(*q).u[a] = b
		(*q).s[b] += (*q).s[a]
	}
	return
}

func (q *qui) quiRootPathComp(a int) int {
	for (*q).u[a] != a {
		n := (*q).u[a]
		// Update below to flatten tree
		(*q).u[a] = (*q).u[n]
		a = n
	}
	return a
}

func (q *qui) QuiConnectedPathComp(a int, b int) bool {
	return q.quiRootPathComp(a) == q.quiRootPathComp(b)
}

func (q *qui) QuiUnionPathComp(a int, b int) {
	a = q.quiRootPathComp(a)
	b = q.quiRootPathComp(b)
	// Root with shorter DFS joins the Larger one
	if (*q).s[a] > (*q).s[b] {
		(*q).u[b] = a
		(*q).s[a] += (*q).s[b]
	} else {
		(*q).u[a] = b
		(*q).s[b] += (*q).s[a]
	}
	return
}

func (u *union) construct(n int) {
	// this implementation can be changed if struct of u changes
	for i := 0; i < n; i++ {
		(*u) = append((*u), i)
	}
}

func (u *union) length() int {
	// this implementation can be modifed if struct of u changes
	return len(*u)
}

func (u *union) root(a int) int {
	for (*u)[a] != a {
		a = (*u)[a]
	}
	return a // (*u)[a]
}

func (u *union) QuickFindUnion(a int, b int) {
	var ia, ib int = (*u)[a], (*u)[b]
	for i := 0; i < u.length(); i++ {
		if (*u)[i] == ia {
			(*u)[i] = ib // update the leader in the component
		}
	}
	return
}

func (u *union) QuickFindConnected(a int, b int) bool {
	return (*u)[a] == (*u)[b]
}

func (u *union) QuickUnionUnion(a int, b int) {
	i := u.root(a)
	j := u.root(b)
	(*u)[i] = j
	return
}

func (u *union) QuickUnionConnected(a int, b int) bool {
	return u.root(a) == u.root(b)
}

// UnionFindMain is the main function
func UnionFindMain() {
	// Take input and test
	var (
		f  QuickFind                 = &union{}
		u  QuickUnion                = &union{}
		qi QUImproved                = &qui{}
		qp QUImprovedPathCompression = &qui{}
	)

	fmt.Println("Testing Quick find")
	f.construct(5)
	f.QuickFindUnion(1, 2)
	f.QuickFindUnion(3, 4)
	f.QuickFindUnion(0, 1)
	fmt.Println(f.QuickFindConnected(0, 2))
	fmt.Println(f.QuickFindConnected(0, 3))
	f.QuickFindUnion(0, 3)
	fmt.Println(f.QuickFindConnected(0, 4))

	fmt.Println("Testing Quick Union")
	u.construct(5)
	u.QuickUnionUnion(1, 2)
	u.QuickUnionUnion(3, 4)
	u.QuickUnionUnion(0, 1)
	fmt.Println(u.QuickUnionConnected(0, 2))
	fmt.Println(u.QuickUnionConnected(0, 3))
	u.QuickUnionUnion(0, 3)
	fmt.Println(u.QuickUnionConnected(0, 4))

	fmt.Println("Testing Quick Union Improved with weights")
	qi.quiConstruct(6)
	qi.QuiUnion(0, 1)
	qi.QuiUnion(2, 3)
	qi.QuiUnion(3, 4)
	qi.QuiUnion(2, 5)
	fmt.Println(qi.QuiConnected(2, 4))
	fmt.Println(qi.QuiConnected(0, 4))
	qi.QuiUnion(1, 4)
	fmt.Println(qi.QuiConnected(3, 1))

	fmt.Println("Testing Quick Union with path comprehension")
	qp.quiConstruct(6)
	qp.QuiUnionPathComp(0, 1)
	qp.QuiUnionPathComp(2, 3)
	qp.QuiUnionPathComp(3, 4)
	qp.QuiUnionPathComp(2, 5)
	fmt.Println(qp.QuiConnectedPathComp(2, 4))
	fmt.Println(qp.QuiConnectedPathComp(0, 4))
	qp.QuiUnionPathComp(1, 4)
	fmt.Println(qp.QuiConnectedPathComp(3, 1))

	fmt.Println("Refer doc for Algorithm: https://github.com/Cbkhare/Algorithms_Robert_Kevin/blob/master/UnionFind.pdf")
}

func main() {
	//fmt.Println("Do something")
	UnionFindMain()
}
