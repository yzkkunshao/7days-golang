package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    int // anonymous field
	innerS     //anonymous field
}

type demo struct {
	a float32
	float32
	string
}

func main() {
	demoInstance := demo{2.1, 33, "hi"}
	fmt.Println("demoInstance is:", demoInstance)
}
