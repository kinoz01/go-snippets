package main

func main() {
	go Code1()
	go Code2()
	go Code3()
	
	select{}
}
