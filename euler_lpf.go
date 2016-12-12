package main

import "fmt"
import "math"

const num int64 = 600851475143

func is_prime(x int64)bool {
    i:= int64(2)
	for ; i < x; i++ {
		if x % i == 0 { return false }
	}
	return true
}

func main() {
    i := int64(math.Sqrt(float64(num)))
	for ; i > 1; i-- {
		if num % i == 0 && is_prime(i) {
			fmt.Println(i)
			break
		}
	}
}
