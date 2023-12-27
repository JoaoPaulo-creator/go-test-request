package main

import (
  "testing"
  "encoding/json"
  "log"
  "bytes"
  "fmt"
  "io"
  "net/http"
)


type PetRequestT struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Category *CategoryRequestT `json:"category"`
}

type CategoryRequestT struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


type ResponseT struct {
  ID int `json:"id"`
  Category CategoryRequestT `json:"category"`
}


func performHTTPRequst(payloadBytes []byte, ch chan<- *http.Response) {
	headers := "application/json"
	request, err := http.Post("https://petstore3.swagger.io/api/v3/pet", headers, bytes.NewBuffer(payloadBytes))
  if err != nil {
    fmt.Println("An error occurred while making the HTTP request")
  }
  ch <- request
}


func TestCasePetCreationWIthChannels(t *testing.T) {
	// montando o payload
	payload := &PetRequestT{
		ID:   10,
		Name: "dog",
		Category: &CategoryRequestT{
			ID:   1,
			Name: "dogie",
		},
	}
	// fazendo o parse da struct para json
	payloadBytes, err := json.Marshal(payload)
  if err != nil {
		t.Fatal("An error occurred while trying to parse the payload")
	}
    
  // criando o canal para a response
  responseChannel := make(chan *http.Response)
  
  // fazendo a chamada
  go performHTTPRequst(payloadBytes, responseChannel)
  
  request := <-responseChannel
  defer request.Body.Close()

  if request.StatusCode != 200 {
		t.Fatal("An error occurred while trying to parse the payload")
	}

	res, err := io.ReadAll(request.Body)
	if err != nil {
		t.Fatalf("An error occurred while reading the response: %d", err)
	}

	var pet ResponseT
	err = json.Unmarshal(res, &pet)

	if err != nil {
		fmt.Println("An error occurred while unmarshalling the JSON: ", err)
	}

	want := 0
	got := pet.ID
	if want != got {
		t.Fatalf("Expected ID %d | Received ID %d", want, got)
	}

	log.Printf("Pet created ID: %d", pet.ID)
	log.Printf("Pet category ID: %d", pet.Category.ID)


}

