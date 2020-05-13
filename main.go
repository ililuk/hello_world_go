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
func sayGoodmorning() {
	fmt.Println("Goodmorning. I am Baba")
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
	var andrey int
	andrey = 500000
	andrey = andrey + 16340
	andrey = andrey * 2
	fmt.Println(andrey)
	var svetlana string
	svetlana = "mom"
	svetlana = "mom" + "i_love_you"
	fmt.Println(svetlana)
	var bogdanAndBee bool
	bogdanAndBee = true
	fmt.Println(bogdanAndBee)
	sayHello()
	sayHello()
}
