package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

func partialMultiply(x int) func(int) int {
	return func(y int) int {
		return multiply(x, y)
	}
}

func add(x, y, z int) int {
	return x + y + z
}

func partialAdd(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return add(x, y, z)
		}
	}
}

func main() {
	double := partialMultiply(2)
	triple := partialMultiply(3)

	fmt.Println(double(6))
	fmt.Println(triple(6))

	add5 := partialAdd(5)
	add2 := add5(2)
	add9 := add2(9)
	fmt.Println(add9)

	fmt.Println(partialAdd(5)(2)(9))

}
