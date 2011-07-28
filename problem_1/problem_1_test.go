package problem_1

import (
	"testing"
	"sort"
	"fmt"
)

func iSliceEq(a, b []int)bool {
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
	return false
}

func TestProvidedMultiples(t *testing.T) {
	correct := []int{3, 5, 6, 9}
	result := multiples([]int{3,5}, 10)
	sort.Ints(result)
	if(!iSliceEq(result, correct)) {
		t.Error("Multiples are incorrect", result, "expecting", correct)
	}
}

func TestProvidedResult(t *testing.T) {
	var correct int = 23
	result := Sum([]int{3,5}, 10)
	if result != correct {
		t.Error("Sum is incorrect", result, "expecting", correct)
	}
}

func TestManualMultiples(t *testing.T) {
	correct := []int{3, 5, 6, 9, 10, 12, 15, 18}
	result := multiples([]int{3,5}, 20)
	sort.Ints(result)
	if(!iSliceEq(result, correct)) {
		t.Error("Multiples are incorrect", result, "expecting", correct)
	}
}

func TestManualResult(t *testing.T) {
	var correct int = 78
	result := Sum([]int{3,5}, 20)
	if result != correct {
		t.Error("Sum is incorrect", result, "expecting", correct)
	}
}

func TestAnswer(t *testing.T) {
	fmt.Println("The answer to problem 1 is", Sum([]int{3,5}, 1000))
}