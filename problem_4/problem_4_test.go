package problem_4

import (
	"testing"
	"fmt"
)

func TestPalindromeChecker(t *testing.T) {
	is_pal := []int{101, 9009, 12321, 555, 1}
	is_not_pal := []int{123, 9008, 12345}
	for _, i := range is_pal {
		if Palindrome(i) != true {
			t.Errorf("Palindrome fail %v is a palindrome", i)
		}
	}
	for _, i := range is_not_pal {
		if Palindrome(i) != false {
			t.Errorf("Palindrome fail %v is not a palindrome", i)
		}
	}
}

func TestExampleData(t *testing.T){
	l := Largest(2)
	if 9009 != l {
		t.Errorf("Largest palindrome product of 2 digits: %v should be 9009", l)
	}
}

func TestResult(t *testing.T){
	fmt.Println("The answer is", Largest(3))
}
