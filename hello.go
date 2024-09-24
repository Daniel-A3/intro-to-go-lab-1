package main

import "fmt"
import "time"

func hello(n int) {
	fmt.Println("Hello from goroutine", n)
}

func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
	}
	time.Sleep(1 * time.Second)
}
