package main

import (
	"fmt"
	"time"
)

func increment(i *int) *int {
	*i = *i + 1
	return i
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	i := 0

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at:", t)
				increment(&i)
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(3100 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}