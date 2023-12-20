package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type PetRequestBody struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Category *Category `json:"category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

func TestShouldCreateDog(t *testing.T) {
	// montando o payload
	payload := &PetRequestBody{
		ID:   10,
		Name: "dog",
		Category: &Category{
			ID:   1,
			Name: "dogie",
		},
	}
	// fazendo o parse da struct para json
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		t.Fatal("An error occurred while trying to parse the payload")
	}

	// setando header e function para a request
	headers := "application/json"
	request, err := http.Post("https://petstore3.swagger.io/api/v3/pet", headers, bytes.NewBuffer(payloadBytes))

	if err != nil {
		t.Fatal("An error occurred while trying to parse the payload")
	}

	defer request.Body.Close()

	if request.StatusCode != 200 {
		t.Fatal("An error occurred while trying to parse the payload")
	}

	fmt.Println(request)

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(request.Body)
	if err != nil {
		fmt.Println("Error reading response body")
	}

	fmt.Println(buffer.String())
}
