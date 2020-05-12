package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello. How are you doing?")
}
func sayGoodbye() {
	fmt.Println("Goodbye. My name is Illia!")
}
func main() {
	sayGoodbye()

	fmt.Println("Hello world.my name is Illia")
	fmt.Println("123")
	var illia string
	illia = "hello"
	illia = "hello" + "world"
	illia = illia + "world"
	fmt.Println(illia)
	var counter int
	counter = counter + 2
	counter = counter + 2
	counter = counter + 2
	fmt.Println(counter)
	sayHello()
	sayHello()
}
