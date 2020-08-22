package main

import (
	"fmt"
)

type list []int

type List interface {
	Scan(n int)
	Print()
	winner()
	fastWinner()
}

func (l *list) Scan(n int) {
	for j := 0; j < n; j++ {
		var k int
		fmt.Scan(&k)
		(*l) = append((*l), k)
	}
	return
}

func (l *list) Print() {
	fmt.Println(*l)
}

func (l *list) winner() {
	h := (*l)[0]
	p := (*l)[1]
	for p > 0 && h > 0 {
		h -= p
		p /= 2
	}
	if h <= 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func (l *list) fastWinner() {
	h := (*l)[0]
	p := (*l)[1]
	if h/2 < p {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

// https://www.codechef.com/AUG20B/problems/CHEFWARS

func main() {
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		// lt := list{} // without interface
		var lt List = &list{}
		lt.Scan(2)
		//lt.Print()
		// lt.winner()
		lt.fastWinner()
	}
}
