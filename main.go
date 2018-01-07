package main

import (
	"net/http"
	"math"
	"log"
)

func main() {
	fileHandler := http.FileServer(http.Dir("static"))
	wrappedHandler := MyPrimeHandler(fileHandler)
	http.ListenAndServe(":80", wrappedHandler)
}

func MyPrimeHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Create some CPU usage
    var n = 2500
    var prime = nthPrime(n)
    log.Printf("%v is the %vth prime number\n", prime, n)
    
    // Serve HTTP
    h.ServeHTTP(w, r)
  })
}

// Calculates the n-th prime
func nthPrime(n int) int {
    foundPrimes := 0
    currentNumber := 0
    for foundPrimes < n {
        if isPrime(currentNumber) {
            foundPrimes++
        }
        currentNumber++
    }
    return currentNumber - 1
}

// Checks, if a given number is prime
func isPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value % i == 0 {
            return false
        }
    }
    return value > 1
}

// Calculates the n-th fibonacci number
func fibonacci(n int) int {
    current, prev := 0, 1
    for i := 0; i < n; i++ {
        current, prev = current + prev, current
    }
    return current
}
