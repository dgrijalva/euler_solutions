package problem_3

import (
	"testing"
	"fmt"
)

func TestPrimeFactorGen(t *testing.T) {
	correct := []int64{5, 7, 13, 29}
	result := Factors(13195)
	if len(result) == len(correct) {
		for i, c := range correct {
			if c != result[i] {
				t.Errorf("Prime factors %v: %v should be %v", i, result[i], c)
			}
		}
	} else {
		t.Errorf("Prime factors %v should be %v", result, correct)
	}
}

func TestAnswer(t *testing.T) {
	fmt.Println("The answer to problem 3 is", MaxFactor(600851475143))
}