package main

func main() {
	go Code1()
	go Code2()
	go Code3()
	go Code4()
	go Code5()
	
	select{}
}
