//NOTE: For this implementation I tried to create my own implementation of The Sieve of Eratosthenes after little research.
// Although not fast enough and more brute force this solution would work.

class Sieve {
  NthPrime(n) {
    let primes = [2, 3, 5, 7];
    // Early return for first 4 primes
    if (n < primes.length) return primes[n];

    let number = 9;

    while (primes.length !== n + 1) {
      // Check division of the primes
      if (!this.CheckDivOfPrimes(number, primes)) {
        primes.push(number);
      }

      // Only check odds as evens are never prime after 2
      number += 2;
    }

    return primes[n];
  }

  CheckDivOfPrimes(n, primes) {
    // Only need to check up to the square root of n
    let limit = Math.ceil(Math.sqrt(n));
    for (let i = 0; i < limit; i++) {
      if (n % primes[i] === 0) {
        return true;
      }
    }

    return false;
  }
}

module.exports = Sieve;
