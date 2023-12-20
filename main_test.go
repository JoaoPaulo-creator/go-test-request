package main

import (
	"net/http"
	"testing"
)

type EqualityChecker interface {
	Equals(other interface{}) bool
}

func equals(a, b interface{}) bool {
	if eq, ok := a.(EqualityChecker); ok {
		return eq.Equals(b)
	}

	return a == b
}

func DoOnRequest[T, A any](got T, want A, t *testing.T) {
	if !equals(got, want) {
		t.Fatalf("got %+v | want: %+v", got, want)
	}
}

func TestShouldGet404(t *testing.T) {
	endpoint := "https://cep.awesomeapi.com.br/json/80010011"
	got, _ := http.Get(endpoint)
	want := 404

	DoOnRequest(got.StatusCode, want, t)
}

func TestShouldGet200(t *testing.T) {
	endpoint := "https://cep.awesomeapi.com.br/json/80010010"
	got, _ := http.Get(endpoint)
	want := 200

	DoOnRequest(got.StatusCode, want, t)
}
