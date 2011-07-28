package problem_2

import (
	"testing"
	"fmt"
)

func TestFibGeneration(t *testing.T) {
	correct := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fib := NewFib()
	for i, c := range correct {
		if v := fib.Next(); v != c {
			t.Errorf("Fibonacci test entry %v is incorrect.  %v should be %v", i, v, c)
		}
	}
	fib.Close()
}

func TestSimpleSum(t *testing.T) {
	correct := (2 + 8 + 34)
	sum := Sum(100)
	if correct != sum {
		t.Errorf("Sum test is incorrect.  %v should be %v", sum, correct)
	}
}

func TestResult(t *testing.T) {
	fmt.Println("The answer to problem 2 is", Sum(4e6))
}