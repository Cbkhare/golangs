package main

import (
	"fmt"
)

type nc struct { // node connection
	prnt string // char
	val  float64
}

type union struct {
	uMap map[string]*nc
	Map  map[string]string
}

type uIntf interface {
	construct(equations *[][]string)
	root(a string) string
	isConnected(a string, b string) bool
	makeUnion(a string, b string, v float64)
	solveQuery(a string, b string) float64
	getVal(a string) float64
}

func (u *union) construct(equations *[][]string) {
	u.Map = make(map[string]string)
	u.uMap = make(map[string]*nc)
	for _, eq := range *equations {
		u.Map[eq[0]] = eq[0]
		u.Map[eq[1]] = eq[1]
	}
	return
}

func (u *union) root(a string) string {
	for u.Map[a] != a {
		a = u.Map[a]
	}
	return a
}

func (u *union) isConnected(a string, b string) bool {
	return u.root(a) == u.root(b)
}

func (u *union) makeUnion(a string, b string, v float64) {
	a1 := u.root(a)
	b1 := u.root(b)

	if a1 != b1 {
		if a1 == a && b1 == b {
			// both are new elements
			u.Map[b] = a
			u.uMap[b] = &nc{
				prnt: a,
				val:  1 / v,
			}
		} else if a1 != a && b1 != b {
			// both have differen parent
			u.Map[b1] = a1
			u.uMap[b1] = &nc{
				prnt: a1,
				val:  1, // both parents would be in the same level
			}
		} else if a1 == a && b1 != b {
			// one of them have parent
			u.Map[a] = b
			u.uMap[a] = &nc{
				prnt: b,
				val:  v,
			}
		} else if a1 != a && b1 == b {
			// one of the havw parent
			u.Map[b] = a
			u.uMap[b] = &nc{
				prnt: a,
				val:  1 / v,
			}
		}
	}
}

func (u *union) getVal(a string) float64 {
	p := 1.0
	for u.Map[a] != a {
		p *= u.uMap[a].val
		//fmt.Println(a,u.Map[a] u.uMap[a].val)
		a = u.Map[a]
	}
	return p
}

func (u *union) solveQuery(a string, b string) float64 {

	if _, ok := u.Map[a]; !ok {
		return -1.0
	} else if _, ok1 := u.Map[b]; !ok1 {
		return -1.0
	} else if !u.isConnected(a, b) {
		return -1.0
	} else if a == b {
		return 1.0
	}
	return u.getVal(a) / u.getVal(b)
}

// https://leetcode.com/problems/evaluate-division/submissions/

func calcEquation(equations [][]string, values []float64, queries [][]string) (ans []float64) {
	m := &union{}
	m.construct(&equations)
	for i, eq := range equations {
		m.makeUnion(eq[0], eq[1], values[i])
	}
	fmt.Println(m.Map)
	for _, q := range queries {
		ans = append(ans, m.solveQuery(q[0], q[1]))
	}
	return
}
