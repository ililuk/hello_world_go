package main

import (
	"fmt"
)

func goodMorning(name string) (string) {
	return "Good morning, " +name
}

func food(name string) (string) {
	return "what's your favorite food " +name +"?"
}

func main() {
	fmt.Println(goodMorning("Illia") + " " +food("Andrei"))
}

	
