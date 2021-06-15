package main

import (
	"testing"
	"fmt"
)

func TestFibTest(t *testing.T) {
	// var a int = 1
	// var b int = 1
	a := 1 //自动类型推断
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println("123")
}
