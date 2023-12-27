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

type Customer struct {
	ID   string
	Name string
}

type Purchase struct {
	ID               string
	Value            float64
	Customer         []*Customer
	DateTimePurchase string
}

func SetPurchase() *Purchase {
	return &Purchase{
		ID:    "034985kosdjf",
		Value: 65456.54,
		Customer: []*Customer{
			{
				ID:   "456456",
				Name: "Ze"},
		},
		DateTimePurchase: "2023-12-21",
	}
}


type PetRequest struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Category *CategoryRequest `json:"category"`
}

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


type Response struct {
  ID int `json:"id"`
  Category CategoryRequest `json:"category"`
}

func TestCasePetCreation(t *testing.T) {
	// montando o payload
	payload := &PetRequest{
		ID:   10,
		Name: "dog",
		Category: &CategoryRequest{
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
 
  // lendo o response da request criada anteriomente
  res, err := io.ReadAll(request.Body) 
  if err != nil {
    t.Fatalf("Ocorreu um erro ao ler response: %d", err)
  }
  // tipando o response
  var pet Response  
  err = json.Unmarshal(res, &pet)

  if err != nil {
    fmt.Println("Ocorreu um erro no unmarshelling do JSON: ", err)
  }
 
  want := 0
  got := pet.ID
  if want != got {
    t.Fatalf("Esperado ID %d | Recebido ID %d", want, got)
  }

  log.Printf("Id do pet criado: %d", pet.ID)  
  log.Printf("Id da categoria do pet: %d", pet.Category.ID)
  
}

