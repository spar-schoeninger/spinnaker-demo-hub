package main

import (
	"net/http"
	"math"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}