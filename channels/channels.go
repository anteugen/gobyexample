package main

import (
    "fmt"
    "time"
)

func pingPong() {
	pingChan := make(chan string)
    pongChan := make(chan string)

    go func() {
        for {
            
            pingChan <- "ping"
            
            msg := <-pongChan
            fmt.Println(msg)
            time.Sleep(time.Second)
        }
    }()

    go func() {
        for {
            msg := <-pingChan
            fmt.Println(msg)
            pongChan <- "pong"
            time.Sleep(time.Second)
        }
    }()

    time.Sleep(5 * time.Second)
}

func worker(ch chan int) {
	outerLoop:
	for i := 100; ; i += 100 {
		select {
		case ch <- i:
			fmt.Println("Sent: ", i)
			time.Sleep(200 * time.Millisecond)
		default:
			fmt.Println("Channel is full")
			break outerLoop
		}
	}
}

func main() {

	pingPong()
    
    ch := make(chan int, 20)

	go worker(ch)

	outerLoop:
	for i := 0; ;i++ {
		select {
		case ch <- i:
			fmt.Println("Sent: ", i)
			time.Sleep(200 * time.Millisecond)
		default:
			fmt.Println("Channel is full")
			break outerLoop
		}
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", <-ch)
	}
	fmt.Println(" ")
}
