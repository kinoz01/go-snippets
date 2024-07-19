package main

func main() {

	go Router0()
	go Router1()
	go Router2()
	go Router3()
	go Router4()
	go Router5()

	select{}
}
