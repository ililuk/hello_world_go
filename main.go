package main

import "fmt"

func hello(abcd string) {
	fmt.Println("Hello,", abcd)
}

func showNumber(one int) {
        fmt.Println(one + 1)
}

func main() {
	var i int
	for i = 0; i < 5; i = i + 1 {
		hello("Illia")
	}
	showNumber(23)
}
