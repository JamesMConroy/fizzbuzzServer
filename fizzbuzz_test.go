package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFizz(t *testing.T) {
	t.Run("returns number", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fizzbuzz/11", nil)
		response := httptest.NewRecorder()

		FizzbuzzServer(response, request)

		got := response.Body.String()
		want := "11"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("returns Buzz", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fizzbuzz/5", nil)
		response := httptest.NewRecorder()

		FizzbuzzServer(response, request)

		got := response.Body.String()
		want := "Buzz"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("returns Fizz", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fizzbuzz/3", nil)
		response := httptest.NewRecorder()

		FizzbuzzServer(response, request)

		got := response.Body.String()
		want := "Fizz"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
	t.Run("returns FizzBuzz", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/fizzbuzz/15", nil)
		response := httptest.NewRecorder()

		FizzbuzzServer(response, request)

		got := response.Body.String()
		want := "FizzBuzz"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
