package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ResultResponse struct {
	Result int `json:"result"`
}

func main() {
	port := "8080"

	log.Printf("Starting HTTP server on port %s", port)

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/v1/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("n"))

		render.JSON(w, r, NewResultResponse(recursiveFibonacci(n)))
	})
	r.Get("/v2/fibonacci", func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("n"))

		render.JSON(w, r, NewResultResponse(memoizedFibonacci(n)))
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func NewResultResponse(result int) *ResultResponse {
	return &ResultResponse{
		Result: result,
	}
}

//recursiveFibonacci calculate the result of fibonacci only with recursion
func recursiveFibonacci(n int) int {
	if n < 2 {
		return n
	}

	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

//memoizedFibonacci calculate the result of fibonacci with recursion and memoization
func memoizedFibonacci(n int) int {
	cache := make(map[int]int)

	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}

		if _, ok := cache[n]; !ok {
			cache[n] = fibonacci(n-1) + fibonacci(n-2)
		}

		return cache[n]
	}

	return fibonacci(n)
}
