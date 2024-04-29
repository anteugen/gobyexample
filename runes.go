package main

import (
	"fmt"
)

func runes(s string) {
	fmt.Printf("%v: ", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#x ", s[i]) 
	}
	fmt.Printf("\n")
}

func main() {
	
	const s = "hello"

	fmt.Println("len: ", len(s))

	runes(s)
	runes("1234")
	runes("สวัสดี")
	runes("漢字")

	for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
}