// random functions that might be useful again later

// get all the primes up to whatever.
// not the most efficient, but quick
func primes(upto int)(primes []int) {
	sieve := make([]bool, upto + 1)
	primes = make([]int, 0)
	for i := 2; i <= upto; i++ {
		if sieve[i] == false {
			primes = append(primes, i)
			for n := i; n < upto; n += i {
				sieve[n] = true
			}
		}
	}
	return
}

func TestPrimeFinding(t *testing.T){
	correct := []int{2, 3, 5, 7, 11, 13, 17, 19}
	result := primes(20)
	for i, c := range correct {
		if result[i] != c {
			t.Errorf("Prime finding error: Entry %v was %v.  Expected %v", i, result[i], c)
		}
	}
}
