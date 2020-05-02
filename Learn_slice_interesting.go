package main

import (
	"fmt"
)

func main() {

	x := [...]int{1, 2, 3, 5, 6}
	s := x[1:3]
	fmt.Println(&s[0], &x[1], &x[2], &x[3])
	s = append(s, 4)
	fmt.Println(s, x, " ", &s[0], &s[1], &s[2])
}

/*
Interesting,
when appended on the slice, slice modifies original array
*/
