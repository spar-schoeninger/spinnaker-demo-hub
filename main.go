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
    var number = 49979687
    var isPrime = IsPrime(number)
    log.Printf("Is %v prime? %v\n", number, isPrime)
    
    // Serve HTTP
    h.ServeHTTP(w, r)
  })
}

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value % i == 0 {
            return false
        }
    }
    return value > 1
}