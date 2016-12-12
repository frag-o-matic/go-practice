package main

import (
	"fmt"
)

func main() {
    var limit=30
	for i := 1; i <= limit; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	if i % 5 == 0 && i % 3 == 0 {
		fmt.Println("FizzBuzz")
	} else if i % 5 == 0 {
		fmt.Println("Buzz")
	} else if i % 3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(i)
	}
}
