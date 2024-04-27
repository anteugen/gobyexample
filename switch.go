package main

import (
	"fmt"
	"math"
)

func isPrime(i int) bool {
	if i == 1 {
		return false
	}

	for n := 2; n <= int(math.Sqrt(float64(i))); n++ {
		if i%n == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := range 15 {
		switch {
		case i == 0:
			fmt.Printf("Number is zero\n")
		case isPrime(i):
			fmt.Printf("Number %v is a prime number\n", i)
			fallthrough
		case i%2 == 0:
			fmt.Printf("Number %v is even\n", i)
		default:
			fmt.Printf("Number %v is odd\n", i)
		}
	}

	score := 75

	fmt.Println("\nFor score of 75:")
	switch {
	case score >= 90:
		fmt.Println("Excellent")
		fallthrough
	case score >= 75:
		fmt.Println("Very good")
		fallthrough
	case score >= 60:
		fmt.Println("Good")
		fallthrough
	case score >= 50:
		fmt.Println("Decent")
		fallthrough
	default:
		fmt.Println("Unsatisfactory")
	}
}