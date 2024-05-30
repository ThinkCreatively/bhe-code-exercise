//NOTE: For this implementation I researched The Sieve of Eratosthenes to greater detial and tried to stay as faithful to it as I could

package sieve

type EtSieve struct {
	primes     []int64
	currentMax int64
}

type Sieve interface {
	NthPrime(n int64) int64
}

func NewSieve() Sieve {
	return &EtSieve{
		primes:     []int64{2, 3, 5, 7},
		currentMax: 7,
	}
}

func (s *EtSieve) NthPrime(n int64) int64 {
	// return early for first 4 primes
	if int64(len(s.primes)) > n {
		return s.primes[n]
	}

	// will be used for generating slice length
	limit := s.currentMax

	for int64(len(s.primes)) <= n {
		// Increase the range of the slice
		limit *= 2

		// Create a sieve slice up to the new limit
		sieveArray := make([]bool, limit+1)

		// Mark non-primes in the range only up to the square root of the limit
		for i := int64(2); i*i <= limit; i++ {
			if !sieveArray[i] {
				for j := i * i; j <= limit; j += i {
					sieveArray[j] = true
				}
			}
		}

		// Collect primes from the sieve array
		for i := s.currentMax + 1; i <= limit; i++ {
			if i >= 2 && !sieveArray[i] {
				s.primes = append(s.primes, i)
			}
		}

		s.currentMax = limit
	}

	return s.primes[n]
}
