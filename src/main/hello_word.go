package main

import (
	"fmt"
	"os"
)

// hello world
func main() {
	fmt.Println(os.Args)
	fmt.Println("hello world")
	os.Exit(-1)
}
