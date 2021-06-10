package fibo

import "testing"

func TestFindFiboWithRecursion(t *testing.T) {
	nInput := int64(10)
	expectedOutput := int64(55)
	realOutput := CalculateWithRecursionMethod(nInput)
	if realOutput != expectedOutput {
		t.Errorf("Got: %v but expected %v", realOutput, expectedOutput)
	}
}

func TestFindFiboWithoutRecursion(t *testing.T) {
	nInput := int64(10)
	expectedOutput := int64(55)
	realOutput := CalculateWithoutRecursiveMethod(nInput)
	if realOutput != expectedOutput {
		t.Errorf("Got: %v but expected %v", realOutput, expectedOutput)
	}
}
