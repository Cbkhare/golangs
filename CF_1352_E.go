package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swapSlice(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func getInputnums() []int {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	strs := strings.Split(line[0:len(line)-1], " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

func specialNums(n *[]int) (count int) {
	count = 0
	mp := make(map[int]bool)
	var s []int
	l := len(*n)
	for i := 1; i < l; i++ {
		l1 := len(s) - 1
		j := 0
		if _, ok := mp[(*n)[i]]; ok {
			count++
		}
		for j < i {
			var x int
			if i-j == 1 {
				x = (*n)[i] + (*n)[i-1]
			} else {
				x = (*n)[i] + s[l1-j]
			}
			s = append(s, x)
			mp[x] = true
			j++
		}
		//fmt.Println(s)
	}
	if count == 0 {
		swapSlice(&(*n))
		//fmt.Println(n)
		return specialNums(&(*n))
	}
	return count
}

func main1352() {
	var a, b int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		n := getInputnums()
		x := specialNums(&n)
		fmt.Println(x)
	}
}
