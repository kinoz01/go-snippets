package main

import "fmt"

func createCounter() func()int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}

func createAdder() func(x int)int {
	var sum int
	return func(x int)int{
		sum +=x
		return sum
	}
}

func outer() func(){
	x := 69
	inner := func () {
		fmt.Println(x) // inner "remembers" x even after outer is done
    }
	return inner
}

func main() {
	fmt.Println("Counting using closure:")
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println("Example when not creatting an instanceof the anonym func:")
	fmt.Println(createCounter()())
	fmt.Println(createCounter()())
	fmt.Println(createCounter()())

	fmt.Println("Adding using closure:")
	add := createAdder()
	fmt.Println(add(5))
	fmt.Println(add(3))
	fmt.Println(add(2))

	fmt.Println("Encapsulation of the variable x:")
	closure := outer()
	closure()

}