package problem_4

// A palindromic number reads the same both ways. The largest 
// palindrome made from the product of two 2-digit numbers is 9009 = 91  99.
// 
// Find the largest palindrome made from the product of two 3-digit numbers.

func rev(i int)(r int) {
	for i > 0 {
		r = (r * 10) + i % 10
		i = i / 10
	}
	return
}

func high(n int)(r int){
	for i := 0; i < n; i++ {
		r = (r * 10) + 9
	}
	return
}

func Palindrome(i int)bool {
	return i == rev(i)
}

func Largest(digits int)(l int) {
	var a, b, c int
	for a = high(digits); a > 0 && a * a > l; a-- {
		for b = a; b > 0 && a * b > l; b-- {
			c = a * b
			if Palindrome(c) {
				l = c
			}
		}
	}
	return
}