package main

import ("fmt")


//func buyFish(size int, price int ) {
//}

//Hello, Andrei. You are 2 years old

func greetingWithAge(name string, age int) {
	fmt.Println("Hello,",name,". You are",age,"years old")
}

func main () {
	fmt.Println("hello world")
	var i int = 0
	for i = 0; i < 6;i = i + 1 {
		greetingWithAge("Andrei",i)	
	}
	}
	
