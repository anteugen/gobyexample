package main

import (
	"fmt"
	"time"
	"context"
	"sync"
)

func timeoutWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()

	c3 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c3 <- "result 3"
    }()

	c4 := make(chan string, 1)
    go func() {
        time.Sleep(4 * time.Second)
        c3 <- "result 4"
    }()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	select {
	case res := <-c4:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {

	var wg sync.WaitGroup

	ctx := context.TODO()
	wg.Add(1)
	go timeoutWithContext(ctx, &wg)

	c1 := make(chan string, 1)
    go func() {
		wg.Add(1)
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
		defer wg.Done()
    }()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		wg.Add(1)
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
		defer wg.Done()
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	wg.Wait()
}