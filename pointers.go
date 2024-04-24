package main

import (
	"fmt"
)

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeropointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)
	zerovalue(i)
	fmt.Println("after zerovalue: ", i)
	zeropointer(&i)
	fmt.Println("after zeropointer: ", i)

	array := [4]int{1, 2, 3, 4}
	zerovalue(array[1])
	fmt.Println("Array element after zerovalue: ", array)
	zeropointer(&array[1])
	fmt.Println("Array element after zerpointer: ", array)
	var pointer *int = &array[2]
	zeropointer(pointer)
	fmt.Println("Array element with assigned var after zeropointer: ", array)
}