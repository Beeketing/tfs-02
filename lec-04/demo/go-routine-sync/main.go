package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := int64(2)
	rwLock := sync.RWMutex{}
	go func() {
		for {
			read(&n, &rwLock)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 5)
	go write(&n, &rwLock)
	time.Sleep(time.Second * 10)
}

func write(n *int64, rwLock *sync.RWMutex) {
	rwLock.Lock()
	defer rwLock.Unlock()
	*n += int64(1000)
	fmt.Println("[Write - Lock] Current time: ", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("[Write - Unlock] Current time: ", time.Now())
}

func read(n *int64, rwLock *sync.RWMutex) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	fmt.Println("[Read - RLock] Current n is: ", *n)
}
