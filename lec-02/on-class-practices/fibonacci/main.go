package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// import local package
	"./fibo"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// read calculating method: recursion or non-recursion
	fmt.Print("Enter the calculating method: ")
	method, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input: %v", err)
		return
	}

	fmt.Print("Enter a number: ")
	textNumber, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input: %v", err)
		return
	}
	// strings.TrimSpace is a method to trim space (space, tab, line break)
	// Eg: "10 " OR "10\n" OR "    10     " will become "10"
	n, err := strconv.Atoi(strings.TrimSpace(textNumber))

	if err != nil {
		fmt.Println("Invalid input: %v", err)
		return
	}

	result := int64(0)
	// use strings.TrimSpace too
	switch strings.TrimSpace(method) {
	case "recursion":
		result = fibo.CalculateWithRecursionMethod(int64(n))
	case "non-recursion":
		result = fibo.CalculateWithoutRecursiveMethod(int64(n))
	default:
		fmt.Println("Not supported method: ", method)
	}
	fmt.Println("Final result: ", result)
}
