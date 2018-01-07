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
    var number = 100
    var fib = fibonacci(number)
    log.Printf("%v is the %vth fibonacci number\n", fib, number)
    
    // Serve HTTP
    h.ServeHTTP(w, r)
  })
}

// Checks, if a given number is prime
func IsPrime(value int) bool {
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
