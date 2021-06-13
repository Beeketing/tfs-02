package main

import "fmt"

func main() {
	shouldRecover := false
	dontPanic(shouldRecover)
	fmt.Println("After calling panic. Your application is not exit when calling panic. It's magical!!!!!!!!")
}

func dontPanic(shouldRecover bool) {
	if shouldRecover {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
	}
	fmt.Println("Calling panic here. You'll see!")
	panic("panic here")
	fmt.Println("This line will not be printed")
}
