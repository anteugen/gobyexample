package main

import (
	"fmt"
)

func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("Total: ", total)
}

type Numerable interface {
	~float32 | ~float64 | int
}

func max[T Numerable](nums ...T) T {
	maxValue := nums[0]

	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}

func customPrintf(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4)
	fmt.Println(max(4, 7, 5, 2))
	fmt.Println(max(21, 36.2, 76.7, 12))
	customPrintf("Today is %v and it's %v degrees outside.\n", "Friday", 13)
}