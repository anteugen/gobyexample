package main

import (
	"fmt"
)

func sequenceGenerator() func() int {
    count := 1
    return func() int {
        count *= 2
        return count
    }
}

func memoizeFibonacci() func(int) int {
    cache := make(map[int]int)
    var fib func(int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        if result, found := cache[n]; found {
            return result
        }
        result := fib(n-1) + fib(n-2)
        cache[n] = result
        return result
    }
    return fib
}

func main() {
	next := sequenceGenerator()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	fib := memoizeFibonacci()
    fmt.Println("Seventh number of the fibonacci sequence:", fib(7))
    fmt.Println("Tenth number of the fibonacci sequence:",fib(10))
}