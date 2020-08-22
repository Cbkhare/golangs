package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const maxInt int = 4294967295

func sortString() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	y := strings.Split(x, "")
	sort.Strings(y)
	fmt.Println(strings.Join(y[:], ""))
}

func mainScan() {
	//fmt.Println("Do something")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	scanner.Scan()
	y := scanner.Text()
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Split(bufio.ScanWords)
	scanner1.Scan()
	z := scanner1.Text()
	fmt.Printf("%T %T %T", x, y, z)
	fmt.Println(x, y, z)
	var t [2]int
	for i := 0; i < 2; i++ {
		fmt.Scanln(&t[i])
	}
	// NOTE Scan or Scaln reads only one non-space separated string
	fmt.Println(t)
}

type list []int

type List interface {
	Scan(n int)
	Print()
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

func scanAndPrint() {
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		// lt := list{} // without interface
		var lt List = &list{}
		lt.Scan(2)
		lt.Print()
	}
}
