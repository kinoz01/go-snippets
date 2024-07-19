package main

import "fmt"


// we try to demonstrate polymorphism using interface and methods

type greeting interface {
	hey()
}

type Person struct {
	Name string
	Age int
}

type Dog struct {
	Name string
	Breed string
}

func (p *Person) hey(){
	fmt.Println("Hello")
}

func (d *Dog) hey(){
	fmt.Println("Ooof oof")
}

func sayHey(welcomer greeting) {
	welcomer.hey()
}

func main() {
	var ayoub *Person
	var scobido Dog

	ayoub.hey()
	scobido.hey()
	
	// using Polymorphism
	sayHey(ayoub)
	sayHey(&scobido)
	
	Test1()
}
