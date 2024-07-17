package main

func main() {

	go Router1()
	go Router2()

	select{}
}
