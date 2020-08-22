package main

import (
	"fmt"
)

const maxInt int = 4294967295

type list []int

func (l *list) Scan(n int) {
	var e int
	for i := 0; i < n; i++ {
		fmt.Scan(&e)
		(*l) = append((*l), e)
	}
}

func (l *list) Print() {
	fmt.Println(*l)
}

func (l *list) minJump(t int) {
	var minVal int
	var jump = maxInt
	for _, v := range *l {
		if t > v && t%v == 0 && t/v < jump {
			jump = t / v
			minVal = v
		}
	}
	if jump == maxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(minVal)
	}
}

//https://www.codechef.com/AUG20B/submit/LINCHESS

func main() {
	var j, l, k int
	fmt.Scan(&j)
	for i := 0; i < j; i++ {
		fmt.Scan(&l)
		fmt.Scan(&k)
		lt := list{}
		lt.Scan(l)
		lt.minJump(k)
	}
}
