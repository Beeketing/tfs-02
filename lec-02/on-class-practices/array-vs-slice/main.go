package main

import "fmt"

func main() {
	// modify slice in another function
	sl := []int64{1, 2, 3, 4}
	fmt.Printf("Slice inside main function (before modifying): %v\n", sl)
	modifySlice(sl)
	fmt.Printf("Slice inside main function (after modifying): %v\n", sl)

	fmt.Println("----------------------------------------")

	// modify array in another function
	arr := [1]int64{99}
	fmt.Printf("Array inside main function (before modifying): %v\n", arr)
	modifyArray(arr)
	fmt.Printf("Array inside main function (after modifying): %v\n", arr)
}

// this is an unexported function
// will modify slice by changing first item value
func modifySlice(sl []int64) {
	if len(sl) == 0 {
		return
	}
	sl[0] = 10000
	fmt.Printf("Slice inside modifySlice function: %v\n", sl)
}

// this is an unexported function
// will modify array by changing first item value
func modifyArray(ar [1]int64) {
	ar[0] = 20000
	fmt.Printf("Array inside modifyArray function: %v\n", ar)
}
