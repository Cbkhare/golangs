package main

import (
	"fmt"
)

// Naming is dummy
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

func mainInterface() {

	x := &name{"foo", "bar"}
	y := test{
		nn: name{"bazinga", "poo"},
	}
	fmt.Println((*x).fname, x.fname, y.nn.fullname())

	// new attmept
	var n Naming
	n = name{fname: "Cb", lname: "k"}
	fmt.Println(n.fullname())

	//fmt.Println(y.nn.fullname())
}
