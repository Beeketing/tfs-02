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
	quit := make(chan bool)

	go generateNumberPair(ch, quit)
	go sum(ch, quit)
	time.Sleep(10 * time.Second)
	// pass 2 bool values
	quit <- true
	quit <- true
	time.Sleep(2 * time.Second)
}

func generateNumberPair(ch chan int32, quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("Quitting generateNumberPair func")
			return
		default:
			n := rand.Int31()
			ch <- n
			// how can it be better?
			time.Sleep(time.Second)
		}
	}
}

func sum(ch chan int32, quit chan bool) {
	sum := int64(0)
	for {
		select {
		case n := <-ch:
			sum += int64(n)
			fmt.Printf("Current sum is %v, n is: %v\n", sum, n)
		case <-quit:
			fmt.Println("Quiting sum func")
			return
		}
	}
}
