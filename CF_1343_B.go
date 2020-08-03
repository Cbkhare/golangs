package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main1343() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b%2 != 0 || (b/2)%2 != 0 {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
			var c []string
			sum := 0
			sum1 := 0
			for j := 0; j < b; j += 2 {
				c = append(c, strconv.Itoa(j+2))
				sum += j + 2
			}
			for j := 1; j < b-2; j += 2 {
				// b-2 to append last value in line 28
				c = append(c, strconv.Itoa(j))
				sum1 += j
			}
			c = append(c, strconv.Itoa(sum-sum1))
			fmt.Println(strings.Join(c, " "))
		}
	}
}

/*
https://codeforces.com/problemset/problem/1343/B
Hint:
sum(half even) == even
sum(half odd) == even, only when the len(odd number)==even
*/
