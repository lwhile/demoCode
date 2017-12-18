package server

func primeNumber(n int) []int {
	primes := make([]int, 0)
	ch := make(chan int)
	go generate(ch)
	//var pass bool
	for i := 0; i < n; i++ {
		prime := <-ch
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
