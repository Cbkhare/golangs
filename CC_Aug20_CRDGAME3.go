package main

import (
	"fmt"
)

func genNewNumLength(n *int) (l int) {
	// var s string
	/*
		for n > 9 {
			// s += "9"
			l++
			n -= 9
		}
		if n != 0 {
			// s += strconv.Itoa(n)
			l++
		}
	*/
	l += (*n) / 9
	if (*n)%9 != 0 {
		l++
	}
	return
}

func numWinner(a *int, b *int) {
	la := genNewNumLength(a)
	lb := genNewNumLength(b)

	if la < lb {
		fmt.Println(0, la)
	} else {
		fmt.Println(1, lb)
	}
}

// https://www.codechef.com/AUG20B/submit/CRDGAME3

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		numWinner(&a, &b)
	}
}
