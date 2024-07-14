package main

import "fmt"

type myFunc func(int, int) int

func addFunc(a, b int) int {
	return a + b
}

func substractFunc(a, b int) int {
	return a - b
}

func multiplyFunc(a int, b int) int {
	return a * b
}

func applyFuncs(x, y int, funcs ...myFunc) {
	for _, opts := range funcs {
		fmt.Println(opts(x, y))
	}
}

func main() {
	applyFuncs(6, 4, addFunc, substractFunc, multiplyFunc)
}
