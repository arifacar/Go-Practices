package main

import (
	"fmt"

)


func main() {
	fmt.Println("JSON Read Test")

	foo1 := new(Foo) // or &Foo{}
	getJson("http://jsonplaceholder.typicode.com/users", foo1)
	println(foo1.Bar)
}