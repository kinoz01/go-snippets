package main

import "fmt"

// Demostrate polymorphism without using methods

type greeting func()

type Person struct {
	Age  int
	Name string
}

type Dog struct {
}

func main() {
	fmt.Println("hello")
}
