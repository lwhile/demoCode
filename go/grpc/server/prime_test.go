package server

import (
	"fmt"
	"testing"
)

func Test_primeNumber(t *testing.T) {
	primes := primeNumber(100)
	fmt.Println(primes)
}
