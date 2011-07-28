package problem_2

// Each new term in the Fibonacci sequence is generated by adding 
// the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// 
// By considering the terms in the Fibonacci sequence whose values 
// do not exceed four million, find the sum of the even-valued terms.

func Sum(max int)(r int) {
	f := NewFib()
	for i := f.Next(); i <= max; i = f.Next() {
		if i & 1 == 0 {
			r += i
		}
	}
	f.Close()
	return
}

// Using a goroutine for this is wildly unnecessary, but what the hell
type Fib struct {
	next chan int
	c chan int
}

func NewFib()*Fib{
	f := new(Fib)
	f.next = make(chan int, 1)
	f.c = make(chan int, 1)
	go f.run()
	return f
}

func (f *Fib) run() {
	var a, b, c int = 0, 1, 0
	for {
		c = a + b
		a, b = b, c
		select {
		case f.next <- c:
		case <- f.c:
			return
		}
	}
}

func (f *Fib) Next()int {
	return <-f.next
}

func (f *Fib) Close() {
	f.c <- 1
}