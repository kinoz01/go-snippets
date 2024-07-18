package main

func main() {

	go Router0()
	go Router1()
	go Router2()

	select{}
}
