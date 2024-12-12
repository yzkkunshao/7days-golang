package main

import (
	"fmt"
)

type Engine struct {
	a string
}

func main() {
	c := &Engine{"haha"}
	t := []*Engine{c}
	fmt.Println(t)
}
