package main

import "fmt"

type myFunc func(int, int) interface{}

func addFunc(a, b int) interface{} {
	return a + b
}

func substractFunc(a, b int) interface{} {
	return a - b
}

func multiplyFunc(a int, b int) interface{} {
	return a * b
}

func divideFunc(a int, b int) interface{} {
	return float64(a) / float64(b)
}

func applyFuncs(x, y int, funcs ...myFunc) {
	for _, opts := range funcs {
		fmt.Println(opts(x, y))
	}
}

func main() {
	applyFuncs(6, 4, addFunc, substractFunc, multiplyFunc, divideFunc)
}
