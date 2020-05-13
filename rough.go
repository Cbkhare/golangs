package main

import (
	"bufio"
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
