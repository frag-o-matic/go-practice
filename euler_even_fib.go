package main

import (
    "fmt"
)

func main() {
    fib:=[2]int{1,2}
    even_fib_sum:=2 //the first even num is 2
    temp:=0
    for temp < 4e6 {
        temp = fib[0] + fib[1]
        fib[0] = fib[1]
        fib[1] = temp

        if temp % 2 == 0 {
            even_fib_sum += temp
        }
    }
    fmt.Println(even_fib_sum)
}
