package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  In this application, we will use go-routine to demonstrate the process of crawling webpages concurrently
*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ch := make(chan int32)
	go generateNumberPair(ch)
	go sum(ch)
	time.Sleep(time.Minute)
}

func generateNumberPair(ch chan int32) {
	for {
		n := rand.Int31()
		ch <- n
		// how can it be better?
		time.Sleep(time.Second)
	}
}

func sum(ch chan int32) {
	sum := int64(0)
	for n := range ch {
		sum += int64(n)
		fmt.Printf("Current sum is %v, n is: %v\n", sum, n)
	}
}
