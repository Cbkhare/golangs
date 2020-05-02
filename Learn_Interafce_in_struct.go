package main

import (
	"fmt"
)

type Naming interface {
	fullname() string
}

type name struct {
	fname string
	lname string
}
  
func (n name) fullname() string {
	return n.fname + "_" + n.lname
}

type test struct {
  // Interface used as a datatype inside struct
	nn Naming
	//nn name
}

func main() {

	x := &name{"foo", "bar"}
	y := test{
		nn: name{"bazinga", "poo"},
	}
	fmt.Println((*x).fname, x.fname, y.nn.fullname())
	//fmt.Println(y.nn.fullname())
}
