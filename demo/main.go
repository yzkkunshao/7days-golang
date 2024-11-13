package main

import (
	"fmt"
)

type myStruct struct{ i int }

var v myStruct
var p *myStruct

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {

	ms := &struct1{10, 15.5, "Joe biden"}

	ms = new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	ms2 := struct1{3, 4, "haha"}

	fmt.Printf("The int is: %d\n", ms2.i1)
	fmt.Printf("The float is: %f\n", ms2.f1)
	fmt.Printf("The string is: %s\n", ms2.str)
	fmt.Println(ms2)
}
