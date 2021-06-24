package fibo

/*
 In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence,
 called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
 starting from 0 and 1.
 Fibonacci Formula: Fn = Fn-1 + Fn-2
 Eg: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……
 With F(0) = 0 and F(1) = 1 Or F(0) = 1 and F(1) = 1
*/

// CalculateWithRecursionMethod calculates fibonacci nth number by using recursion method
func CalculateWithRecursionMethod(n int64) int64 {
	// If n <= 1, just returns 1
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	// recursively returns the same fibonacci function with n-1 and n-2 based on Fibonacci Formula
	return CalculateWithRecursionMethod(n-1) + CalculateWithRecursionMethod(n-2)
}

// CalculateWithoutRecursiveMethod calculates fibonacci nth number without using recursion
// We must use a loop inside the function replace recursion method
func CalculateWithoutRecursiveMethod(n int64) int64 {
	// If n <= 1, just returns immediately for fast calculation
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}

	// calculate nth number from beginning of fibonacci number string by using Formula:
	// Fn = Fn-1 + Fn-2
	nth := int64(0)
	first, second := int64(0), int64(1)
	for i := int64(2); i <= n; i++ {
		nth = first + second
		first, second = second, nth
	}
	return nth
}
