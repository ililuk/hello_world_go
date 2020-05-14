package main

import "fmt"

func helloGrandfather() {
	fmt.Println("hello grandfather")
}

func main() {
	var i int
	for i = 0; i < 5; i = i + 1 {
		fmt.Println(i)
		helloGrandfather()
	}
}
