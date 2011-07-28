package problem_1

// If we list all the natural numbers below 10 that are multiples 
// of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// 
// Find the sum of all the multiples of 3 or 5 below 1000.

func multiples(of []int, below int)[]int {
	// Using a hash as a poor man's set
	values := make(map[int]bool)
	for _, num := range of {
		for i := num; i < below; i += num {
			values[i] = true
		}
	}
	results := make([]int, len(values))
	var i int = 0
	for n, _ := range values {
		results[i] = n
		i++
	}
	return results
}

func Sum(multiples_of []int, below int)(result int) {
	for _, i := range multiples(multiples_of, below) {
		result += i
	}
	return
}