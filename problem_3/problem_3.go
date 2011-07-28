package problem_3

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600,851,475,143 ?

func Factors(value int64)(r []int64) {
	r = make([]int64, 0)
	for i := int64(2); i <= value; i++ {
		if value % i == 0 {
			value = value / i
			r = append(r, i)
		}
	}
	return
}

func MaxFactor(value int64)int64 {
	f := Factors(value)
	return f[len(f) - 1]
}