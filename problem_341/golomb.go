package golomb

type Entry struct {
	N uint64
	GN uint64
}

func Sum(max uint64)uint64{
	gen := StartGenerator()
	var accum uint64
	var next_n uint64 = 1
	var next_n3 uint64 = 1
	for {
		entry := <-gen
		if entry.N == next_n3 {
			accum += entry.GN
			next_n++
			next_n3 = next_n * next_n * next_n
		}
		if entry.N == max {
			break
		}
	}
	return accum
}

func StartGenerator()(c chan Entry) {
	gen := make(chan Entry, 1024)
	go startGenerator(gen)
	return gen
}

func startGenerator(c chan Entry) {
	var i uint64
	for {
		i++
		c <- Entry{i,1}
	}
}